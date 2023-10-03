package walletManager

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"

	"goEthTestTask/crypto"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Wallet struct {
	accounts.Wallet
}

type WalletManager struct {
	keyStore     *keystore.KeyStore
	passFilePath string
	cipherer     crypto.Cipherer
}

func NewWalletManager(keystorePath string, passFilePath string, aesKey string) WalletManager {
	return WalletManager{
		keyStore:     keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP),
		passFilePath: passFilePath,
		cipherer:     crypto.NewCipherer(aesKey),
	}
}

func (wm WalletManager) WalletExist() bool {
	return len(wm.keyStore.Wallets()) != 0
}

func (wm WalletManager) CreateWallet(password string) (Wallet, error) {
	if wm.WalletExist() {
		return Wallet{}, errors.New("wallet already exist")
	}

	hashString := crypto.Sha256Hash(password)
	encryptedPassword, err := wm.cipherer.EncryptAES([]byte(hashString))

	err = os.WriteFile(wm.passFilePath, encryptedPassword, 0644)
	if err != nil {
		return Wallet{}, err
	}

	_, err = wm.keyStore.NewAccount(password)
	if err != nil {
		return Wallet{}, err
	}

	wallet := Wallet{wm.keyStore.Wallets()[0]}

	err = wm.keyStore.Unlock(wallet.Accounts()[0], password)
	if err != nil {
		return Wallet{}, err
	}
	return wallet, nil
}

func (wm WalletManager) WalletLogin(password string) (Wallet, error) {

	if !wm.WalletExist() {
		return Wallet{}, errors.New("wallet not exist")
	}

	address := wm.keyStore.Accounts()[0]

	encryptedFileData, err := os.ReadFile(wm.passFilePath)
	if err != nil {
		return Wallet{}, err
	}

	decryptedData, err := wm.cipherer.DecryptAES(encryptedFileData)

	if string(decryptedData) == crypto.Sha256Hash(password) {
		err = wm.keyStore.Unlock(address, password)
		if err != nil {
			return Wallet{}, err
		}
		fmt.Println("Wallet unlocked")
		return Wallet{
			wm.keyStore.Wallets()[0],
		}, nil
	}

	return Wallet{}, errors.New("wrong password")

}

func (wm WalletManager) Lock(wallet Wallet) error {
	if wallet == (Wallet{}) {
		return errors.New("wallet not defined")
	}

	if !wm.WalletExist() {
		return errors.New("wallet not exist")
	}
	err := wm.keyStore.Lock(wallet.Accounts()[0].Address)
	if err != nil {
		return err
	}
	return nil
}

func (wm WalletManager) SendTestTx(wallet Wallet) error {
	if wallet == (Wallet{}) {
		return errors.New("wallet not defined")
	}
	if !wm.WalletExist() {
		return errors.New("wallet not exist")
	}

	client, err := ethclient.Dial("https://goerli.blockpi.network/v1/rpc/public")
	if err != nil {
		return err
	}
	nonce, err := client.PendingNonceAt(context.Background(), wallet.Accounts()[0].Address)
	if err != nil {
		return err
	}

	toAddress := wallet.Accounts()[0].Address
	value := big.NewInt(1_000_000_000_000_000_000)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	var data []byte

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	sTx, err := wallet.Wallet.SignTx(wallet.Accounts()[0], tx, big.NewInt(5))
	if err != nil {
		return err
	}
	err = client.SendTransaction(context.Background(), sTx)
	if err != nil {
		return err
	}
	return nil
}
