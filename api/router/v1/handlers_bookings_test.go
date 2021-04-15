package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eimlav/go-gym/testutils"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func generateBookingsPOSTRequest(memberName string, classEventID int) (req BookingsPOSTRequest, err error) {
	req.MemberName = memberName
	req.ClassEventID = &classEventID

	return req, nil
}

func TestHandlersBookingsPOST(t *testing.T) {
	sqlDB, mock, err := testutils.SetupMock()
	assert.NoError(t, err)
	defer sqlDB.Close()

	testRequest, err := generateBookingsPOSTRequest("Bob", 1)
	if err != nil {
		assert.Error(t, err)
	}

	rows := sqlmock.NewRows(
		[]string{"id", "created_at", "updated_at", "deleted_at", "class_id", "date"},
	).AddRow("1", time.Now(), time.Now(), nil, "1", time.Now())
	mock.ExpectQuery("SELECT \\* FROM `class_events`").WithArgs(1).WillReturnRows(rows)
	mock.ExpectExec("INSERT INTO `bookings`").WillReturnResult(sqlmock.NewResult(1, 1))

	w := httptest.NewRecorder()
	router := gin.Default()
	router.POST("/api/v1/bookings", HandleBookingsPOST)

	testRequestJSON, _ := json.Marshal(testRequest)
	req, _ := http.NewRequest("POST", "/api/v1/bookings", bytes.NewBuffer(testRequestJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestHandlersBookingsPOST_ClassEventDoesNotExist(t *testing.T) {
	sqlDB, mock, err := testutils.SetupMock()
	assert.NoError(t, err)
	defer sqlDB.Close()

	testRequest, err := generateBookingsPOSTRequest("Bob", 1)
	if err != nil {
		assert.Error(t, err)
	}

	mock.ExpectQuery("SELECT \\* FROM `class_events`").WithArgs(1).WillReturnRows(&sqlmock.Rows{})

	w := httptest.NewRecorder()
	router := gin.Default()
	router.POST("/api/v1/bookings", HandleBookingsPOST)

	testRequestJSON, _ := json.Marshal(testRequest)
	req, _ := http.NewRequest("POST", "/api/v1/bookings", bytes.NewBuffer(testRequestJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestHandlersBookingsPOST_InvalidParameters(t *testing.T) {
	testData := []struct {
		MemberName   string
		ClassEventID int
	}{
		{"a", 1},
		{"", 0},
	}

	testCases := []BookingsPOSTRequest{}
	for _, test := range testData {
		testRequest, err := generateBookingsPOSTRequest(test.MemberName, test.ClassEventID)
		if err != nil {
			assert.Error(t, err)
		}
		testCases = append(testCases, testRequest)
	}

	w := httptest.NewRecorder()
	router := gin.Default()
	router.POST("/api/v1/bookings", HandleBookingsPOST)

	for _, test := range testCases {
		testRequestJSON, _ := json.Marshal(test)
		req, _ := http.NewRequest("POST", "/api/v1/bookings", bytes.NewBuffer(testRequestJSON))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	}
}
