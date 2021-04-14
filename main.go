package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/eimlav/go-gym/api"
)

func awaitOSSignal() {
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt)
	<-done
}

func main() {
	apiServer := api.NewAPIServer()

	log.Printf("Starting API server on %s", apiServer.GetAddress())
	go apiServer.Start()

	awaitOSSignal()

	log.Println("Shutting down API server.")
	apiServer.Shutdown()
}
