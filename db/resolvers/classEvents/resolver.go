package classEvents

import (
	"github.com/eimlav/go-gym/db/models"
	"gorm.io/gorm"
)

// CreateClassEvent creates a ClassEvent record.
func CreateClassEvent(db *gorm.DB, classEvent *models.ClassEvent) error {
	return db.Create(&classEvent).Error
}

// Exists checks for the existence of a ClassEvent.
func Exists(db *gorm.DB, id uint) (bool, error) {
	classEvent := &models.ClassEvent{}
	err := db.Where("id = ?", id).Find(classEvent).Error
	if err != nil {
		return false, err
	}

	return classEvent.ID != 0, nil
}
