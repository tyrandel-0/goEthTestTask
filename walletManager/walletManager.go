package walletManager

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"goEthTestTask/crypto"
	"log"
	"os"
)

type Wallet struct {
	Account  accounts.Account
	keystore *keystore.KeyStore
}

func CreateWallet(keyStore *keystore.KeyStore, passFilePath string, cipherer crypto.Cipherer) (Wallet, error) {
	var password string
	fmt.Print("Придумайте пароль: ")
	_, err := fmt.Scanf("%s", &password)
	if err != nil {
		return Wallet{}, err
	}

	hashString := crypto.Sha256Hash(password)
	encryptedPassword, err := cipherer.EncryptAES([]byte(hashString))

	err = os.WriteFile(passFilePath, encryptedPassword, 0644)
	if err != nil {
		return Wallet{}, err
	}

	// Создаем новый аккаунт
	account, err := keyStore.NewAccount(password)
	if err != nil {
		return Wallet{}, err
	}

	return Wallet{
		account, keyStore,
	}, nil
}

func WalletLogin(keyStore *keystore.KeyStore, passFilePath string, cipherer crypto.Cipherer) (Wallet, error) {
	address := keyStore.Accounts()[0]

	var password string
	fmt.Print("Введите пароль от адресса " + address.Address.String() + ": ")
	_, err := fmt.Scanf("%s", &password)
	if err != nil {
		return Wallet{}, err
	}

	encryptedFileData, err := os.ReadFile(passFilePath)
	if err != nil {
		return Wallet{}, err
	}

	decryptedData, err := cipherer.DecryptAES(encryptedFileData)

	if string(decryptedData) == crypto.Sha256Hash(password) {
		_, err := keyStore.Export(address, password, password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("кошелек разблокирован")
		return Wallet{
			keyStore.Accounts()[0], keyStore,
		}, nil
	} else {
		fmt.Println("Неверный пароль")
		return Wallet{}, errors.New("wrong password")
	}
}
