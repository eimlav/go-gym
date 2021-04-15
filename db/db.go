package db

import (
	"github.com/eimlav/go-gym/db/models"
	"gorm.io/gorm"

	_ "gorm.io/driver/sqlite"
)

var DB *gorm.DB

// GetDB gets the database instance.
func GetDB() *gorm.DB {
	return DB
}

// SetupDatabase creates a new database instance.
func SetupDatabase(db *gorm.DB) error {
	DB = db

	return nil
}

// Run GORM AutoMigrate function on database.
func MigrateDatabase(db *gorm.DB) error {
	return db.AutoMigrate(models.Class{}, models.ClassEvent{})
}
