package classes

import (
	"github.com/eimlav/go-gym/db/models"
	"gorm.io/gorm"
)

func CreateClass(db *gorm.DB, class *models.Class) error {
	return db.Create(&class).Error
}
