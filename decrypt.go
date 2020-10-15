package hmm

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func Decrypt(keyStr, encryptedStr string) (string, error) {
	key := []byte(keyStr)
	encryptedBytes := []byte(encryptedStr)
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := encryptedBytes[:nonceSize], encryptedBytes[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", nil
	}
	return string(plaintext), nil
}
