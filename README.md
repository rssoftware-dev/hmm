# Hmm

Hmm is a tool used for encrypting strings. One reason to do this is to encrypt config values. A scenario that comes to mind is a key that unfortunately has to be shared between envs and you may not want everyone in a team having access to this.

The public pkg includes 2 main functions Encrypt and Decrypt.

Decrypt is what should be used inside of your apps to decrypt the config values.

## Usage

encrypt will return a base64 encoded string of the encrypted value. NO we don't just wrap it in base64 :D.

```sh
hmm encrypt areallylongpasswordisagoodpasswd "some data"
```

```sh
hmm decrypt areallylongpasswordisagoodpasswd "sZwyys0qgn6peOghyErF7HR_4eF77YWZ6w=="
```
