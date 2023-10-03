package utils

import (
	"fmt"
	"os"

	"github.com/c-bata/go-prompt"
	"golang.org/x/crypto/ssh/terminal"
)

func Completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "create wallet", Description: "create new wallet"},
		{Text: "login to wallet", Description: "login to existing wallet"},
		{Text: "send test Tx", Description: "sends test transaction"},
		{Text: "lock wallet", Description: "locks existing wallet"},
		{Text: "exit", Description: "exit from program"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func GetPassword() string {
	fmt.Print("Password: ")
	bytePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return ""
	}
	fmt.Println()
	return string(bytePassword)
}
