package classes

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eimlav/go-gym/db"
	"github.com/eimlav/go-gym/db/models"
	"github.com/eimlav/go-gym/testutils"
	"github.com/stretchr/testify/assert"
)

func TestCreateClass(t *testing.T) {
	sqlDB, mock, err := testutils.SetupMock()
	assert.NoError(t, err)
	defer sqlDB.Close()
	gormDB := db.GetDB()

	mock.ExpectExec("INSERT INTO `classes`").WillReturnResult(sqlmock.NewResult(1, 1))

	startDate, err := time.Parse(time.RFC3339, "2021-04-20T15:00:00+00:00")
	assert.NoError(t, err)
	endDate, err := time.Parse(time.RFC3339, "2021-04-22T15:00:00+00:00")
	assert.NoError(t, err)

	class := &models.Class{
		Name:      "Yoga",
		StartDate: startDate,
		EndDate:   endDate,
		Capacity:  uint(10),
	}

	err = CreateClass(gormDB, class)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), class.ID)
}
