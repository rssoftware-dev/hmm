package secrets_test

import (
	"testing"

	"github.com/rssoftware-dev/hmm/secrets"
)

func TestEncryption(t *testing.T) {
	key := "passphrasewhichneedstobe32bytes!"
	dataToEncrypt := "my encrypted data"
	t.Run("Test encrypt and decrypt", func(t *testing.T) {
		encrypted, err := secrets.Encrypt(key, dataToEncrypt)
		if err != nil {
			t.Fatal(err)
		}
		decrypted, err := secrets.Decrypt(key, encrypted)
		if err != nil {
			t.Fatal(err)
		}
		if decrypted != dataToEncrypt {
			t.Fatalf("%s != %s", decrypted, dataToEncrypt)
		}
	})
}
