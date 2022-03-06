package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fileutils"
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Print("Enter key : ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)

	key, _ := hex.DecodeString(input)

	for _, file := range fileutils.GetAllFiles() {
		hextext, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err.Error())
		}

		text, _ := hex.DecodeString(string(hextext))

		block, err := aes.NewCipher(key)
		if err != nil {
			panic(err.Error())
		}

		aesGCM, err := cipher.NewGCM(block)
		if err != nil {
			panic(err.Error())
		}

		nonceSize := aesGCM.NonceSize()
		nonce, ciphertext := text[:nonceSize], text[nonceSize:]

		plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			panic(err.Error())
		}

		err2 := ioutil.WriteFile(file, plaintext, 0644)

		if err2 != nil {
			panic(err2.Error())
		}
	}

}
