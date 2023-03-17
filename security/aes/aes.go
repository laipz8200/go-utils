package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type AES struct {
	key []byte
}

func (a *AES) Load(key []byte) error {
	if len(key) != 32 {
		return errors.New("key must be 32 bytes")
	}
	a.key = key
	return nil
}

func (a *AES) Encrypt(plaintext []byte) (string, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plaintext))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(cipherText), nil
}

func (a *AES) Decrypt(cipherText string) ([]byte, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}

	cipherTextBytes, err := base64.URLEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}

	if len(cipherTextBytes) < aes.BlockSize {
		return nil, errors.New("cipher text too short")
	}

	iv := cipherTextBytes[:aes.BlockSize]
	cipherTextBytes = cipherTextBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherTextBytes, cipherTextBytes)

	return cipherTextBytes, nil
}
