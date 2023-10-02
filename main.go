package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"goEthTestTask/crypto"
	"goEthTestTask/walletManager"
	"log"
	"os"
)

const (
	keystorePath     = "./keystore"         // Путь к директории с keystore
	passwordHashFile = "./passwordHashFile" // Путь к файлу с хешем пароля
	aesKey           = "passphrasewhichneedstobe32bytes!"
)

func main() {
	files, err := os.ReadDir(keystorePath)
	if err != nil {
		log.Fatal(err)
	}
	cipherer := crypto.NewCipherer(aesKey)
	keyStore := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)

	if len(files) == 0 {
		wallet, err := walletManager.CreateWallet(keyStore, passwordHashFile, cipherer)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(wallet.Account.Address.String())
	} else {
		wallet, err := walletManager.WalletLogin(keyStore, passwordHashFile, cipherer)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(wallet.Account.Address.String())
	}
}
