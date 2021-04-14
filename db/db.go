package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// GetDB gets the database instance.
func GetDB() *gorm.DB {
	return DB
}

// SetupDatabase creates a new database instance.
func SetupDatabase() error {
	db, err := gorm.Open(sqlite.Open("go-gym.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	if err = migrateDatabase(db); err != nil {
		return err
	}

	DB = db

	return nil
}

// Run GORM AutoMigrate function on database.
func migrateDatabase(db *gorm.DB) error {
	return db.AutoMigrate(Class{}, ClassEvent{})
}
