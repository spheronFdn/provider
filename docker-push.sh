make docker-image
docker tag ghcr.io/akash-network/provider:latest-amd64 spheronnetwork/devnet-provider:latest-amd64
docker tag ghcr.io/akash-network/provider:latest-arm64 spheronnetwork/devnet-provider:latest-arm64
docker push spheronnetwork/devnet-provider:latest-amd64
docker push spheronnetwork/devnet-provider:latest-arm64