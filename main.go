package main

import (
	"flag"
	"fmt"
	"net/http"
	"github.com/tmilewski/goenv"
	"os"
)

var (
	NWorkers = flag.Int("n", 4, "The number of workers to start")
	HTTPAddr = flag.String("http", "127.0.0.1:8000", "Address to listen for HTTP requests on")
)

func init(){
	err:=goenv.Load()
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func main() {
	// Start the dispatcher.
	StartDispatcher(*NWorkers)

	// Run migrations
	RunMigrations()

	// Register http handler
	http.HandleFunc("/mpesa", Collector)

	// Start Server
	fmt.Println("HTTP server listening on", *HTTPAddr)
	if err := http.ListenAndServe(*HTTPAddr, nil); err != nil {
		fmt.Println(err.Error())
	}
}
