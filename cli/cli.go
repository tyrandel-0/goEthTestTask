package cli

import (
	"fmt"

	"goEthTestTask/cli/handlers"
	"goEthTestTask/cli/utils"
	"goEthTestTask/walletManager"

	"github.com/c-bata/go-prompt"
)

func Options(manager walletManager.WalletManager) {
	fmt.Println("Select option")
	var wallet walletManager.Wallet
	var err error
	for {
		command := prompt.Input("> ", utils.Completer)

		switch command {
		case "create wallet":
			tmpWallet, err := handlers.HandleCreation(manager)
			if err != nil {
				fmt.Println("Error: " + err.Error())
			} else {
				wallet = tmpWallet
			}
		case "login to wallet":
			tmpWallet, err := handlers.HandleLogin(manager)
			if err != nil {
				fmt.Println("Error: " + err.Error())
			} else {
				wallet = tmpWallet
			}
		case "send test Tx":
			err = handlers.HandleTestTx(wallet, manager)
			if err != nil {
				fmt.Println("Error: " + err.Error())
			}
		case "lock wallet":
			err = handlers.HandleLock(&wallet, manager)
			if err != nil {
				fmt.Println("Error: " + err.Error())
			}
		case "exit":
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}
