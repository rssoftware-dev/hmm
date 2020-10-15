package hmm

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// Encrypt will encrypt the given data with the given key using aes.
// The key must be 32 bytes long, example: "passphrasewhichneedstobe32bytes!"
//
// returns the encrypted string or an error.
func Encrypt(key, data string) (string, error) {
	text := []byte(data)
	k := []byte(key)

	c, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encryptedText := gcm.Seal(nonce, nonce, text, nil)
	return string(encryptedText), err
}
