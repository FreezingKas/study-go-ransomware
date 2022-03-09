package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cretz/bine/tor"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
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
		writeDatabase(key)

	default:
		fmt.Fprintf(w, "Tssss")
	}
}

func writeDatabase(key string) { // Add error management

	db, err := sql.Open("sqlite3", "./key.db")

	checkErr(err)

	statement, err := db.Prepare("INSERT INTO keys(timestamp, key) values(?,?)")
	checkErr(err)

	statement.Exec(fmt.Sprint(time.Now().Unix()), key)
	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
