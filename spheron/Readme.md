### Required Dependencies

```shell
brew update
brew upgrade
brew tap ethereum/ethereum
brew install solidity
brew install protoc
brew install protobuf
brew tap ethereum/ethereum
brew install ethereum
```

```shell
cd spheron
solc --bin --abi contract/requestLogger.sol -o build
abigen  --bin=build/RequestLogger.bin --abi=build/RequestLogger.abi --pkg=gen --out=gen/requestLogger.go
```
