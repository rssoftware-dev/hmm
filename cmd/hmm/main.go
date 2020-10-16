package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rssoftware/hmm"
)

func main() {
	args := os.Args[1:]
	switch args[0] {
	case "encrypt":
		if len(args) == 1 {
			fmt.Println("missing key")
			return
		}
		if len(args) == 2 {
			fmt.Println("missing data to encrypt")
			return
		}
		encryptedData, err := hmm.Encrypt(args[1], args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(encryptedData)
	case "decrypt":
	}
}
