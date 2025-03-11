
BUILD_DATE=`date +%Y%m%d%H%M%S`
BUILD_COMMIT=`git describe --tags --always`

abcore:
	@echo "Building ABCore version..."
	go build -ldflags "-X github.com/ABFoundationGlobal/MultiSignatureWallet/cli.buildCommit=${BUILD_COMMIT}\
	    -X github.com/ABFoundationGlobal/MultiSignatureWallet/cli.buildDate=${BUILD_DATE}"

abiot:
	@echo "Modifying go.mod for ABIoT version..."
	cp go.mod go.mod.bak
	echo "replace github.com/ethereum/go-ethereum => github.com/ABFoundationGlobal/abiot v1.9.18-ab-1.3" >> go.mod
	go mod tidy
	go build -ldflags "-X github.com/ABFoundationGlobal/MultiSignatureWallet/cli.buildCommit=${BUILD_COMMIT}\
    	    -X github.com/ABFoundationGlobal/MultiSignatureWallet/utils.buildDate=${BUILD_DATE}" -o MultiSignatureWallet-ABIoT
	mv go.mod.bak go.mod
	go mod tidy

all: abcore abiot
