package main

import (
	"fmt"
	"goEthTestTask/cli"

	"goEthTestTask/walletManager"
)

const (
	keystorePath     = "./storage/keystore"
	passwordHashFile = "./storage/passwordHashFile"
	aesKey           = "passphrasewhichneedstobe32bytes!"
)

func main() {
	manager := walletManager.NewWalletManager(keystorePath, passwordHashFile, aesKey)

	fmt.Printf("\n\n")

	if manager.WalletExist() {
		fmt.Println("Wallet found, use 'login to wallet' to use it")
	} else {
		fmt.Println("Wallet not created, use 'create wallet' to create wallet")
	}

	cli.Options(manager)

}
