rm -rf build
solc --bin --abi protocol-contracts/contracts/NodeProviderRegistery.sol -o ./build
mkdir -p gen/NodeProviderRegistry
abigen  --bin=build/NodeProviderRegistry.bin --abi=build/NodeProviderRegistry.abi --pkg=NodeProviderRegistry --out=gen/NodeProviderRegistry/NodeProviderRegistery.go
mkdir -p gen/requestLogger
solc --bin --abi  protocol-contracts/contracts/requestLogger.sol -o build
abigen  --bin=build/RequestLogger.bin --abi=build/RequestLogger.abi --pkg=requestLogger --out=gen/requestLogger/requestLogger.go
mkdir -p gen/OrderMatching
solc --bin --abi  protocol-contracts/contracts/OrderMatching.sol -o build
abigen  --bin=build/OrderMatching.bin --abi=build/OrderMatching.abi --pkg=OrderMatching --out=gen/OrderMatching/OrderMatching.go
mkdir -p gen/TokenRegistry
solc --bin --abi  protocol-contracts/contracts/TokenRegistery.sol -o build
abigen  --bin=build/TokenRegistry.bin --abi=build/TokenRegistry.abi --pkg=TokenRegistry --out=gen/TokenRegistry/TokenRegistry.go
