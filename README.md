# Hmm

Hmm is tool used by use for config encryption. This allows us to add configs to git knowing they are encrypted before the push happens. 

The idea is to encrypt the config value using the cli tool and then add that to your config file. 

The public pkg includes 2 main functions Encrypt and Decrypt. 

Decrypt is what should be used inside of your apps to decrypt the config values. 