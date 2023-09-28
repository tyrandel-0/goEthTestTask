package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	files, err := os.ReadDir("./keystore")
	if err != nil {
		log.Fatal(err)
	}

	// Создаем новый Keystore
	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)

	if len(files) == 0 {

		// Генерируем новый приватный ключ
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}

		var password string
		fmt.Print("Придумайте пароль: ")
		_, err = fmt.Scanf("%s", &password)
		if err != nil {
			return
		}

		// Создаем новый аккаунт
		account, err := ks.ImportECDSA(privateKey, password)
		if err != nil {
			log.Fatal(err)
		}

		// Получаем адрес созданного аккаунта
		address := account.Address.Hex()

		// Выводим информацию о созданном кошельке
		fmt.Println("Адрес кошелька:", address)
	} else {
		address := ks.Accounts()[0]

		var password string
		fmt.Print("Введите пароль от адресса " + address.Address.String() + ": ")
		_, err = fmt.Scanf("%s", &password)
		if err != nil {
			return
		}

		prkey, err := ks.Export(address, password, password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(prkey))
	}
}
