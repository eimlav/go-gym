package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/eimlav/go-gym/db"

	"github.com/eimlav/go-gym/api"
)

func awaitOSSignal() {
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt)
	<-done
}

func main() {
	// Setup database and API server
	if err := db.SetupDatabase(); err != nil {
		panic(err.Error())
	}

	apiServer, err := api.NewAPIServer()
	if err != nil {
		panic(err.Error())
	}

	log.Printf("API server listening on %s", apiServer.GetAddress())
	go apiServer.Start()

	// Await an exit signal from the OS
	awaitOSSignal()

	// Gracefully shut down the API server.
	log.Println("Shutting down API server.")
	apiServer.Shutdown()
}
