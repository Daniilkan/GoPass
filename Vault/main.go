package Vault

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
)

func Encrypt(plainText string, key []byte) (string, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.New("invalid key size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	plainTextBytes := []byte(plainText)
	padding := blockSize - len(plainTextBytes)%blockSize
	padText := append(plainTextBytes, bytes.Repeat([]byte{byte(padding)}, padding)...)

	iv := make([]byte, blockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	cipherText := make([]byte, len(padText))
	mode.CryptBlocks(cipherText, padText)

	result := append(iv, cipherText...)

	return base64.StdEncoding.EncodeToString(result), nil
}

func Decrypt(encryptedText string, key []byte) (string, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.New("invalid key size")
	}

	cipherText, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	if len(cipherText) < blockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := cipherText[:blockSize]

	cipherText = cipherText[blockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)

	plainText := make([]byte, len(cipherText))
	mode.CryptBlocks(plainText, cipherText)

	padding := int(plainText[len(plainText)-1])
	if padding > len(plainText) {
		return "", errors.New("invalid padding")
	}

	plainText = plainText[:len(plainText)-padding]

	return string(plainText), nil
}
