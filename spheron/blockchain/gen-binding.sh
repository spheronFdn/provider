rm -rf build
solc --bin --abi protocol-contracts/contracts/NodeProviderRegistery.sol -o build --overwrite
mkdir -p gen/NodeProviderRegistry
abigen  --bin=build/NodeProviderRegistry.bin --abi=build/NodeProviderRegistry.abi --pkg=NodeProviderRegistry --out=gen/NodeProviderRegistry/NodeProviderRegistery.go
mkdir -p gen/OrderMatching
solc --bin --abi  protocol-contracts/contracts/OrderMatching.sol -o build --overwrite
abigen  --bin=build/OrderMatching.bin --abi=build/OrderMatching.abi --pkg=OrderMatching --out=gen/OrderMatching/OrderMatching.go
mkdir -p gen/TokenRegistry
solc --bin --abi  protocol-contracts/contracts/TokenRegistery.sol -o build --overwrite
abigen  --bin=build/TokenRegistry.bin --abi=build/TokenRegistry.abi --pkg=TokenRegistry --out=gen/TokenRegistry/TokenRegistry.go
