package testutils

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eimlav/go-gym/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupMock for testing.
func SetupMock() (sqlDB *sql.DB, mock sqlmock.Sqlmock, err error) {
	sqlDB, mock, err = sqlmock.New()

	if err != nil {
		return nil, nil, fmt.Errorf("error setting up mock: %v", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("error connecting to database: %v", err)
	}

	if err := db.SetupDatabase(gormDB); err != nil {
		return nil, nil, fmt.Errorf("error setting up database: %v", err)
	}

	return sqlDB, mock, nil
}

// MockTime is used to mock a time value for sql-mock.
type MockTime struct {
	Time string
}

// Match satisfies sqlmock.Argument interface
func (m MockTime) Match(v driver.Value) bool {
	value, ok := v.(time.Time)

	actualTime, err := time.Parse(time.RFC3339, m.Time)
	if err != nil {
		return false
	}

	return ok && value.UTC() == actualTime.UTC()
}
