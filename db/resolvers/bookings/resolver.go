package bookings

import (
	"github.com/eimlav/go-gym/db/models"
	"gorm.io/gorm"
)

func CreateBooking(db *gorm.DB, booking *models.Booking) error {
	return db.Create(&booking).Error
}
