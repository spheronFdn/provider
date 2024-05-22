#!/bin/bash
# WARNING: the runtime of this script should NOT exceed 5 seconds! (Perhaps can be amended via AKASH_BID_PRICE_SCRIPT_PROCESS_TIMEOUT env variable)
# Requirements:
# curl jq bc mawk ca-certificates
set -o pipefail

function map_token {
  local token=$1

  if [ "$token" == "USDTT" ]; then
    echo "USDT"
  else
    echo "$token"
  fi
}

function get_token_price {
  # cache token price for 60 minutes to reduce the API pressure as well as to slightly accelerate the bidding (+5s)
  CACHE_FILE=/tmp/tokenprice.cache
  if ! test $(find $CACHE_FILE -mmin -60 2>/dev/null); then
    
    mapped_token=$(map_token "$1")
    
    ## cache expired
    usd_per_token=$(curl -s --connect-timeout 3 --max-time 3 -X GET "https://api-osmosis.imperator.co/tokens/v2/price/$mapped_token" -H 'accept: application/json' | jq -r '.price' 2>/dev/null)

    # update the cache only when API returns a result.
    # this way provider will always keep bidding even if API temporarily breaks (unless pod gets restarted which will clear the cache)
    if [ ! -z $usd_per_token ]; then
      # check price is an integer/floating number
      re='^[0-9]+([.][0-9]+)?$'
      if ! [[ $usd_per_token =~ $re ]]; then
        echo -n "$usd_per_token is not an integer/floating number!" >&2
        exit 1
      fi

      # make sure price is in the permitted range
      if ! (( $(echo "$usd_per_token > 0" | bc -l) && \
              $(echo "$usd_per_token <= 1000000" | bc -l) )); then
        echo -n "$usd_per_token is outside the permitted range (>0, <=1000000)" >&2
        exit 1
      fi

      echo "$usd_per_token" > $CACHE_FILE
    fi
  fi

  # Fail if the script can't read CACHE_FILE for some reason
  set -e
  usd_per_token=$(cat $CACHE_FILE)
  echo $usd_per_token
  set +e
}

# bid script starts reading the deployment order request specs here (passed by the Akash Provider)
data_in=$(jq .)

# Pull the pricing data from the deployment request
hasPrice=$(echo "$data_in" | jq -r 'has("price")?')

# default price precision to 12 (for backward compatibility)
precision=$(jq -r '.price_precision? // 12' <<<"$data_in")

# If the price parameter is set, new rate calculations will be used
# otherwise, the original rate calculations will be used (for backward compatibility)
if [[ "$hasPrice" == true ]]; then
  isObject=$(jq -r 'if .price?|type == "object" then true else false end' <<<"$data_in")
  if [[ "$isObject" != true ]]; then
    echo -n "price must be an object! make sure you are using the latest akash-provider." >&2
    exit 1
  fi
  denom=$(jq -r '.price.denom' <<<"$data_in")
  amount=$(jq -r '.price.amount' <<<"$data_in")

  # strip off the .price by setting data_in to .resources
  data_in=$(echo "$data_in" | jq -r '.resources')
fi

# Calculate the resources requested (CPU, memory, storage, IPs, endpoints, GPUs)
##
cpu_requested=$(echo "$data_in" | jq -r '(map(.cpu * .count) | add) / 1000')
memory_requested=$(echo "$data_in" | jq -r '(map(.memory * .count) | add) / pow(1024; 3)' | awk '{printf "%.12f\n", $0}')
ephemeral_storage_requested=$(echo "$data_in" | jq -r '[.[] | (.storage[] | select(.class == "ephemeral").size // 0) * .count] | add / pow(1024; 3)' | awk '{printf "%.12f\n", $0}')
hdd_pers_storage_requested=$(echo "$data_in" | jq -r '[.[] | (.storage[] | select(.class == "beta1").size // 0) * .count] | add / pow(1024; 3)' | awk '{printf "%.12f\n", $0}')
ssd_pers_storage_requested=$(echo "$data_in" | jq -r '[.[] | (.storage[] | select(.class == "beta2").size // 0) * .count] | add / pow(1024; 3)' | awk '{printf "%.12f\n", $0}')
nvme_pers_storage_requested=$(echo "$data_in" | jq -r '[.[] | (.storage[] | select(.class == "beta3").size // 0) * .count] | add / pow(1024; 3)' | awk '{printf "%.12f\n", $0}')
ips_requested=$(echo "$data_in" | jq -r '(map(.ip_lease_quantity//0 * .count) | add)')
endpoints_requested=$(echo "$data_in" | jq -r '(map(.endpoint_quantity//0 * .count) | add)')

# Provider sets the Price he wants to charge in USD/month
##
TARGET_CPU="${PRICE_TARGET_CPU:-1.60}"                   # USD/thread-month
TARGET_MEMORY="${PRICE_TARGET_MEMORY:-0.80}"             # USD/GB-month
TARGET_HD_EPHEMERAL="${PRICE_TARGET_HD_EPHEMERAL:-0.02}" # USD/GB-month
TARGET_HD_PERS_HDD="${PRICE_TARGET_HD_PERS_HDD:-0.01}"   # USD/GB-month (beta1)
TARGET_HD_PERS_SSD="${PRICE_TARGET_HD_PERS_SSD:-0.03}"   # USD/GB-month (beta2)
TARGET_HD_PERS_NVME="${PRICE_TARGET_HD_PERS_NVME:-0.04}" # USD/GB-month (beta3)
TARGET_ENDPOINT="${PRICE_TARGET_ENDPOINT:-0.05}"         # USD for port/month
TARGET_IP="${PRICE_TARGET_IP:-5}"                        # USD for leased IP/month

# GPU pricing per GPU model (USD/GPU unit a month) calculation
##

# Populate the price target gpu_mappings dynamically based on the "price_target_gpu_mappings" value passed by the helm-chart
declare -A gpu_mappings=()

IFS=',' read -ra PAIRS <<< "${PRICE_TARGET_GPU_MAPPINGS}"
for pair in "${PAIRS[@]}"; do
  IFS='=' read -ra KV <<< "$pair"
  key="${KV[0]}"
  value="${KV[1]}"
  gpu_mappings["$key"]=$value
done

# Default to 100 USD/GPU per unit a month when PRICE_TARGET_GPU_MAPPINGS is not set
# Or use the highest price from PRICE_TARGET_GPU_MAPPINGS when model detection fails (ref. https://github.com/akash-network/support/issues/139 )
gpu_unit_max_price=100
for value in "${gpu_mappings[@]}"; do
  # Hint: bc <<< "$a > $b" (if a is greater than b, it will return 1, otherwise 0)
  if bc <<< "$value > $gpu_unit_max_price" | grep -qw 1; then
    gpu_unit_max_price=$value
  fi
done

if ! [[ -z $DEBUG_BID_SCRIPT ]]; then
  echo "DEBUG: gpu_unit_max_price $gpu_unit_max_price"
fi

gpu_price_total=0
while IFS= read -r resource; do
  count=$(echo "$resource" | jq -r '.count')
  model=$(echo "$resource" | jq -r '.gpu.attributes.vendor | (.nvidia // .amd // empty).model // 0')
  vram=$(echo "$resource" | jq -r --arg v_model "$model" '.gpu.attributes.vendor | (
      .nvidia | select(.model == $v_model) //
      .amd | select(.model == $v_model) //
      empty
  ).ram // 0')
  interface=$(echo "$resource" | jq -r --arg v_model "$model" '.gpu.attributes.vendor | (
      .nvidia | select(.model == $v_model) //
      .amd | select(.model == $v_model) //
      empty
  ).interface // 0')
  gpu_units=$(echo "$resource" | jq -r '.gpu.units // 0')
  # default to 100 USD/GPU per unit a month when PRICE_TARGET_GPU_MAPPINGS is not set
  # GPU <vram> price_target_gpu_mappings can specify <model.vram> or <model>. E.g. a100.40Gi=900,a100.80Gi=1000 or a100=950
  if [[ "$vram" != "0" ]]; then
    model="${model}.${vram}"
  fi
  # GPU <interface>: price_target_gpu_mappings can specify <model.vram.interface> or <model.interface>. E.g. a100.80Gi.pcie=900,a100.pcie=1000 or a100.80Gi.sxm4,a100.sxm4 or a100=950
  if [[ "$interface" != "0" ]]; then
    model="${model}.${interface}"
  fi

  # Fallback logic to find the best matching price if vram/interface weren't set in PRICE_TARGET_GPU_MAPPINGS
  if [[ -n "${gpu_mappings["$model"]}" ]]; then
    price="${gpu_mappings["$model"]}"
  elif [[ -n "${gpu_mappings["${model%.*}"]}" ]]; then  # Remove the interface or vram if it's not found
    price="${gpu_mappings["${model%.*}"]}"
  elif [[ -n "${gpu_mappings["${model%%.*}"]}" ]]; then  # Remove vram (and interface if exists)
    price="${gpu_mappings["${model%%.*}"]}"
  else
    price="$gpu_unit_max_price"  # Default catchall price
  fi
  gpu_price_total=$(bc -l <<< "$gpu_price_total + ($count * $gpu_units * $price)")

  if ! [[ -z $DEBUG_BID_SCRIPT ]]; then
    echo "DEBUG: model $model"
    echo "DEBUG: price for this model $price"
    echo "DEBUG: gpu_units $gpu_units"
    echo "DEBUG: gpu_price_total $gpu_price_total"
    echo "DEBUG: count $count"
  fi
done <<< "$(echo "$data_in" | jq -rc '.[]')"

# Calculate the total resource cost for the deployment request in USD
##
total_cost_usd_target=$(bc -l <<< "( \
  ($cpu_requested * $TARGET_CPU) + \
  ($memory_requested * $TARGET_MEMORY) + \
  ($ephemeral_storage_requested * $TARGET_HD_EPHEMERAL) + \
  ($hdd_pers_storage_requested * $TARGET_HD_PERS_HDD) + \
  ($ssd_pers_storage_requested * $TARGET_HD_PERS_SSD) + \
  ($nvme_pers_storage_requested * $TARGET_HD_PERS_NVME) + \
  ($endpoints_requested * $TARGET_ENDPOINT) + \
  ($ips_requested * $TARGET_IP) + \
  ($gpu_price_total) \
  )")


if ! [[ -z $DEBUG_BID_SCRIPT ]]; then
  echo "DEBUG: Total cost USD/month: $total_cost_usd_target"
fi

# average block time: 0.25 seconds
# average number of days in a month: 30.437
# (60/0.25)*60*24*30.437 = 10519027 blocks per month
# Convert the total resource cost for the deployment request into the rate based on the given token or default to USD/block rate

blocks_a_month=10519027

if [[ "$hasPrice" = true ]]; then
  usd_per_token=$(get_token_price $denom)
  total_cost_token=$(bc -l <<<"(${total_cost_usd_target}/$usd_per_token)")
  rate_per_block_token=$(bc -l <<<"(${total_cost_token}/${blocks_a_month})")
  printf "%.*f" "$precision" "$rate_per_block_token"
else
  rate_per_block_usd=$(bc -l <<<"(${total_cost_usd_target}/${blocks_a_month})")
  total_cost_usd="$(printf "%.*f" $precision $rate_per_block_usd)"
  printf "%.f" "$total_cost_usd"
fi