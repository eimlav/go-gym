package classEvents

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eimlav/go-gym/db"
	"github.com/eimlav/go-gym/db/models"
	"github.com/eimlav/go-gym/testutils"
	"github.com/stretchr/testify/assert"
)

func TestCreateClassEvent(t *testing.T) {
	sqlDB, mock, err := testutils.SetupMock()
	assert.NoError(t, err)
	defer sqlDB.Close()
	gormDB := db.GetDB()

	mock.ExpectExec("INSERT INTO `class_events`").WillReturnResult(sqlmock.NewResult(1, 1))

	date, err := time.Parse(time.RFC3339, "2021-04-20T15:00:00+00:00")
	assert.NoError(t, err)

	classEvent := &models.ClassEvent{
		ClassID: uint(1),
		Date:    date,
	}

	err = CreateClassEvent(gormDB, classEvent)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), classEvent.ID)
}
