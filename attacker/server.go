package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)
func hello(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		var key string
		for dict, element := range r.URL.Query() {
			fmt.Println("Key:", dict, "=>", "Value:", element[0])
			key = element[0]
			break
		}
		fmt.Println(key)
		f, err := os.OpenFile("key.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		if _, err := f.WriteString(key+"\n"); err != nil {
			log.Println(err)
		}
		defer f.Close()
		


	default:
		fmt.Fprintf(w, "Tssss")
	}
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}