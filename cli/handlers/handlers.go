package handlers

import (
	"goEthTestTask/cli/utils"
	"goEthTestTask/walletManager"
)

func HandleCreation(manager walletManager.WalletManager) (walletManager.Wallet, error) {
	wallet, err := manager.CreateWallet(utils.GetPassword())
	if err != nil {
		return walletManager.Wallet{}, err
	}
	return wallet, nil
}

func HandleLogin(manager walletManager.WalletManager) (walletManager.Wallet, error) {
	wallet, err := manager.WalletLogin(utils.GetPassword())
	if err != nil {
		return walletManager.Wallet{}, err
	}
	return wallet, nil
}

func HandleTestTx(wallet walletManager.Wallet, manager walletManager.WalletManager) error {
	err := manager.SendTestTx(wallet)
	if err != nil {
		return err
	}
	return nil
}

func HandleLock(wallet *walletManager.Wallet, manager walletManager.WalletManager) error {
	err := manager.Lock(*wallet)
	if err != nil {
		return err
	}
	return nil
}
