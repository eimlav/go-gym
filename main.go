package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/eimlav/go-gym/config"

	"github.com/eimlav/go-gym/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/eimlav/go-gym/api"
)

func awaitOSSignal() {
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt)
	<-done
}

func connectDatabase() error {
	// Setup database and API server
	gormDB, err := gorm.Open(sqlite.Open("go-gym.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.SetupDatabase(gormDB); err != nil {
		return err
	}

	if err := db.MigrateDatabase(db.GetDB()); err != nil {
		return err
	}

	return nil
}

func main() {
	// Get configuration
	if err := config.GetConfig(); err != nil {
		panic(err.Error())
	}

	// Setup database and API server
	if err := connectDatabase(); err != nil {
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
