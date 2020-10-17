package secrets

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// Encrypt will encrypt the given data with the given key using aes.
// The key must be 32 bytes long, example: "passphrasewhichneedstobe32bytes!"
//
// returns the encrypted string or an error.
func Encrypt(key, data string) (string, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	encryptedText := make([]byte, aes.BlockSize+len(data))
	iv := encryptedText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(c, iv)
	stream.XORKeyStream(encryptedText[aes.BlockSize:], []byte(data))
	return base64.URLEncoding.EncodeToString(encryptedText), nil
}
