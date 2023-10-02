package crypto

import (
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
)

func Sha256Hash(inputString string) string {
	hasher := sha256.New()
	hasher.Write([]byte(inputString))
	hash := hasher.Sum(nil)

	hashString := hex.EncodeToString(hash)
	return hashString
}

func encryptAESBlock(data []byte, key []byte) ([]byte, error) {

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	out := make([]byte, len(data))

	cipher.Encrypt(out, data)

	return out, nil
}

func decryptAESBlock(encryptedData []byte, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	out := make([]byte, len(encryptedData))
	cipher.Decrypt(out, encryptedData)

	return out, nil
}

func splitDataToBlocks(text []byte, blockSize uint) [][]byte {
	bs := int(blockSize)
	numBlocks := (len(text) + bs - 1) / bs
	blocks := make([][]byte, numBlocks)

	for i := 0; i < numBlocks; i++ {
		start := i * bs
		end := (i + 1) * bs
		if end > len(text) {
			end = len(text)
		}
		block := make([]byte, bs) // Создаем блок нулей длиной blockSize
		copy(block, text[start:end])
		blocks[i] = block
	}

	return blocks
}
