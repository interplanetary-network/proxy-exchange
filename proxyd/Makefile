GO := go
BINARY_NAME := proxyd

all:
	$(GO) build -trimpath -o $(BINARY_NAME)

darwin_amd64:
	GOOS=darwin GOARCH=amd64 $(GO) build -trimpath -o $(BINARY_NAME)_darwin_amd64

darwin_arm64:
	GOOS=darwin GOARCH=arm64 $(GO) build -trimpath -o $(BINARY_NAME)_darwin_arm64

darwin_universal: darwin_amd64 darwin_arm64
	lipo $(BINARY_NAME)_darwin_amd64 $(BINARY_NAME)_darwin_arm64 -create -output $(BINARY_NAME)_darwin_universal 

linux_amd64:
	GOOS=linux GOARCH=amd64 $(GO) build -trimpath -o $(BINARY_NAME)_linux_amd64

linux_arm64:
	GOOS=linux GOARCH=arm64 $(GO) build -trimpath -o $(BINARY_NAME)_linux_arm64

windows_amd64:
	GOOS=windows GOARCH=amd64 $(GO) build -trimpath -o $(BINARY_NAME)_windows_amd64

contracts:
	make -C ../contracts abi
	test -d internal/exchange || mkdir internal/exchange
	abigen --abi ../contracts/build/abi/ProxyExchange.abi --pkg exchange --out internal/exchange/exchange.go

