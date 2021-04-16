package bookings

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eimlav/go-gym/db"
	"github.com/eimlav/go-gym/db/models"
	"github.com/eimlav/go-gym/testutils"
	"github.com/stretchr/testify/assert"
)

func TestCreateBooking(t *testing.T) {
	sqlDB, mock, err := testutils.SetupMock()
	assert.NoError(t, err)
	defer sqlDB.Close()
	gormDB := db.GetDB()

	mock.ExpectExec("INSERT INTO `bookings`").WillReturnResult(sqlmock.NewResult(1, 1))

	booking := &models.Booking{
		MemberName:   "Bob",
		ClassEventID: 1,
	}

	err = CreateBooking(gormDB, booking)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), booking.ID)
}
