package main

import (
	"github.com/ABFoundationGlobal/MultiSignatureWallet/cli"
)

//go:generate go run chains/generate.go

func main() {
	cli.NewCLI().Execute()
}
