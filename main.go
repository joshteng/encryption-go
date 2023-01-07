package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"

	"encryption/core"
)

func getPassword() []byte {
	fmt.Println("Input key used to encrypt: ")
	password, err := terminal.ReadPassword(0)
	if err != nil {
		log.Fatal(err)
	}
	return password
}

func encrypt() {
	password := getPassword()

	fmt.Println("Input message to encrypt:")
	data, err := terminal.ReadPassword(0)

	if err != nil {
		log.Fatal(err)
	}

	ciphertext, err := core.Encrypt(password, data)
	if err != nil {
		log.Fatal(err)
	}

	sciphertext := hex.EncodeToString(ciphertext)
	fmt.Printf("ciphertext: %s\n", sciphertext)
}

func decrypt() {
	password := getPassword()

	fmt.Println("Input message to decrypt:")
	data, err := terminal.ReadPassword(0)

	if err != nil {
		log.Fatal(err)
	}

	ciphertext, err := hex.DecodeString(string(data))
	if err != nil {
		log.Fatal(err)
	}

	plaintext, err := core.Decrypt(password, ciphertext)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("plaintext: %s\n", plaintext)

}

func main() {
	fmt.Println("Key in 1 to encrypt, 2 to decrypt:")
	reader := bufio.NewReader(os.Stdin)
	key, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	key = strings.TrimSuffix(key, "\n")

	if key == "1" {
		fmt.Println("Encryption mode")
		encrypt()
	} else {
		fmt.Println("Decryption mode")
		decrypt()
	}
}
