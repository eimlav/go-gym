package classEvents

import (
	"github.com/eimlav/go-gym/db/models"
	"gorm.io/gorm"
)

func CreateClassEvent(db *gorm.DB, classEvent *models.ClassEvent) error {
	return db.Create(&classEvent).Error
}
