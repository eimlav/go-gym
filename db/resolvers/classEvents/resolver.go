package classEvents

import (
	"github.com/eimlav/go-gym/db/models"
	"github.com/go-errors/errors"
	"gorm.io/gorm"
)

// CreateClassEvent creates a ClassEvent record.
func CreateClassEvent(db *gorm.DB, classEvent *models.ClassEvent) error {
	return db.Create(&classEvent).Error
}

// Exists checks for the existence of a ClassEvent.
func Exists(db *gorm.DB, id uint) (bool, error) {
	err := db.Where("id = ?", id).Find(&models.ClassEvent{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
