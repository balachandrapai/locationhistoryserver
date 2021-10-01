package main

import (
	"fmt"
	"locationhistoryserver/cmd"
	"locationhistoryserver/db"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// dbService
var dbService *db.Service

// signal handler
func signalHandler() {
	var captureSignal = make(chan os.Signal, 1)
	signal.Notify(captureSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)
	fmt.Println("\nCaught signal: ", <-captureSignal)
	os.Exit(0)
}

// handlerFunc
func handlerFunc(rw http.ResponseWriter, req *http.Request) {
	// Check the type of the http method
	if req.Method == http.MethodGet {
		cmd.RetrieveLocationHandler(rw, req, dbService)
	}
	if req.Method == http.MethodDelete {
		cmd.DeleteLocationHandler(rw, req, dbService)
	}
	if req.Method == http.MethodPost {
		cmd.AppendLocationHandler(rw, req, dbService)
	}
}

// start HTTP server
func server() {
	fmt.Println("Starting server on port 8080...")
	http.HandleFunc("/location/", handlerFunc)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func main() {
	// capture signals
	go signalHandler()

	// Start the DB service
	dbService = db.NewService()

	// start HTTP server
	server()
}
