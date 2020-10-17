package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rssoftware-dev/hmm"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("----------------")
		fmt.Println("Hmm by RS Software.")
		fmt.Println("github.com/rssoftware-dev/hmm")
		fmt.Println("----------------")

		fmt.Println("Missing command.")
		fmt.Println("Avaliable commands: encrypt & decrypt.")
		return
	}

	args := os.Args[1:]

	if len(args) == 1 {
		fmt.Println("missing key")
		return
	}
	if len(args) == 2 {
		fmt.Println("missing data")
		return
	}

	switch args[0] {
	case "encrypt":
		data, err := hmm.Encrypt(args[1], args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(data)
		return
	case "decrypt":
		data, err := hmm.Decrypt(args[1], args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(data)
		return
	}
}
