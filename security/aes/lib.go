package aes

var std = &AES{}

func Load(key []byte) error {
	return std.Load(key)
}

func Encrypt(plaintext []byte) (string, error) {
	return std.Encrypt(plaintext)
}

func Decrypt(cipherText string) ([]byte, error) {
	return std.Decrypt(cipherText)
}
