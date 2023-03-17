package aes

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	aes := AES{}
	aes.Load([]byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
	plaintext := []byte("Hello, world!")
	cipherText, err := aes.Encrypt(plaintext)
	if err != nil {
		t.Errorf("Error encrypting: %v", err)
	}
	decrypted, err := aes.Decrypt(cipherText)
	if err != nil {
		t.Errorf("Error decrypting: %v", err)
	}
	if string(decrypted) != string(plaintext) {
		t.Errorf("Expected %s, but got %s", plaintext, decrypted)
	}
}

func TestDecryptInvalidCipherText(t *testing.T) {
	aes := AES{}
	aes.Load([]byte("my-secret-key"))
	_, err := aes.Decrypt("invalid-cipher-text")
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

func TestDecryptShortCipherText(t *testing.T) {
	aes := AES{}
	aes.Load([]byte("my-secret-key"))
	_, err := aes.Decrypt("MTIzNDU2Nzg5MA==")
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

func TestLoad(t *testing.T) {
	aes := AES{}
	err := aes.Load([]byte("kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk"))
	if err != nil {
		t.Errorf("Error loading key: %v", err)
	}
}

func TestLoadInvalidKey(t *testing.T) {
	aes := AES{}
	err := aes.Load([]byte("invalid-key"))
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}
