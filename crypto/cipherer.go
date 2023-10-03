package crypto

import (
	"crypto/aes"
)

type Cipherer struct {
	key []byte
}

// NewCipherer returns new Cipherer
func NewCipherer(key string) Cipherer {
	return Cipherer{
		key: []byte(key),
	}
}

func (c Cipherer) EncryptAES(data []byte) ([]byte, error) {
	out := make([]byte, 0)
	blocks := splitDataToBlocks(data, aes.BlockSize)
	for _, block := range blocks {
		encryptedBlock, err := encryptAESBlock(block, c.key)
		if err != nil {
			return nil, err
		}
		out = append(out, encryptedBlock...)
	}
	return out, nil
}

func (c Cipherer) DecryptAES(encryptedData []byte) ([]byte, error) {
	out := make([]byte, 0)
	blocks := splitDataToBlocks(encryptedData, aes.BlockSize)
	for _, block := range blocks {
		decryptedBlock, err := decryptAESBlock(block, c.key)
		if err != nil {
			return nil, err
		}
		out = append(out, decryptedBlock...)
	}
	return out, nil
}
