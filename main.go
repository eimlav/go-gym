package main

import (
	"os"
	"os/signal"

	"github.com/eimlav/go-gym/config"

	log "github.com/sirupsen/logrus"

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
		log.Fatal(err)
	}

	// Setup database and API server
	if err := connectDatabase(); err != nil {
		log.Fatal(err)
	}

	apiServer, err := api.NewAPIServer()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("API server listening on %s", apiServer.GetAddress())
	go apiServer.Start()

	// Await an exit signal from the OS
	awaitOSSignal()

	// Gracefully shut down the API server.
	log.Println("Shutting down API server.")
	apiServer.Shutdown()
}
