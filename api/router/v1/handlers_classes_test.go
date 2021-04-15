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

func generateClassesPOSTRequest(name, startDate, endDate string, capacity int) (req ClassesPOSTRequest, err error) {
	if startDate != "" {
		testStartDate, err := time.Parse(time.RFC3339, startDate)
		if err != nil {
			return ClassesPOSTRequest{}, err
		}
		req.StartDate = testStartDate
	}

	if endDate != "" {
		testEndDate, err := time.Parse(time.RFC3339, endDate)
		if err != nil {
			return ClassesPOSTRequest{}, err
		}
		req.EndDate = testEndDate
	}

	req.Name = name
	req.Capacity = &capacity

	return req, nil
}

func TestHandlersClassesPOST(t *testing.T) {
	sqlDB, mock, err := testutils.SetupMock()
	assert.NoError(t, err)
	defer sqlDB.Close()

	testRequest, err := generateClassesPOSTRequest("My Class", "2021-04-20T15:00:00+00:00", "2021-04-20T15:00:00+00:00", 10)
	if err != nil {
		assert.Error(t, err)
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `classes`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO `class_events`").WithArgs(
		sqlmock.AnyArg(), sqlmock.AnyArg(), nil, 1, testutils.MockTime{Time: "2021-04-20T15:00:00+00:00"},
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	w := httptest.NewRecorder()
	router := gin.Default()
	router.POST("/api/v1/classes", HandleClassesPOST)

	testRequestJSON, _ := json.Marshal(testRequest)
	req, _ := http.NewRequest("POST", "/api/v1/classes", bytes.NewBuffer(testRequestJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestHandlersClassesPOST_MultipleEvents(t *testing.T) {
	sqlDB, mock, err := testutils.SetupMock()
	assert.NoError(t, err)
	defer sqlDB.Close()

	testRequest, err := generateClassesPOSTRequest("My Class", "2021-04-20T15:00:00+00:00", "2021-04-22T15:00:00+00:00", 10)
	if err != nil {
		assert.Error(t, err)
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `classes`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO `class_events`").WithArgs(
		sqlmock.AnyArg(), sqlmock.AnyArg(), nil, 1, testutils.MockTime{Time: "2021-04-20T15:00:00+00:00"},
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO `class_events`").WithArgs(
		sqlmock.AnyArg(), sqlmock.AnyArg(), nil, 1, testutils.MockTime{Time: "2021-04-21T15:00:00+00:00"},
	).WillReturnResult(sqlmock.NewResult(2, 1))
	mock.ExpectExec("INSERT INTO `class_events`").WithArgs(
		sqlmock.AnyArg(), sqlmock.AnyArg(), nil, 1, testutils.MockTime{Time: "2021-04-22T15:00:00+00:00"},
	).WillReturnResult(sqlmock.NewResult(3, 1))
	mock.ExpectCommit()

	w := httptest.NewRecorder()
	router := gin.Default()
	router.POST("/api/v1/classes", HandleClassesPOST)

	testRequestJSON, _ := json.Marshal(testRequest)
	req, _ := http.NewRequest("POST", "/api/v1/classes", bytes.NewBuffer(testRequestJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestHandlersClassesPOST_InvalidParameters(t *testing.T) {
	testData := []struct {
		Name      string
		StartDate string
		EndDate   string
		Capacity  int
	}{
		{"My Class", "2021-04-20T15:00:00+00:00", "2021-04-18T15:00:00+00:00", 10},
		{"My Class", "2021-04-20T15:00:00+00:00", "2021-04-20T15:00:00+00:00", 10},
		{"My Class", "2021-04-22T15:00:00+00:00", "2021-04-18T15:00:00+00:00", 10},
		{"My Class", "2021-04-20T15:00:00+00:00", "2021-04-18T15:00:00+00:00", 0},
		{"", "2021-04-20T15:00:00+00:00", "2021-04-18T15:00:00+00:00", 10},
	}

	testCases := []ClassesPOSTRequest{}
	for _, test := range testData {
		testRequest, err := generateClassesPOSTRequest(test.Name, test.StartDate, test.EndDate, test.Capacity)
		if err != nil {
			assert.Error(t, err)
		}
		testCases = append(testCases, testRequest)
	}

	w := httptest.NewRecorder()
	router := gin.Default()
	router.POST("/api/v1/classes", HandleClassesPOST)

	for _, test := range testCases {
		testRequestJSON, _ := json.Marshal(test)
		req, _ := http.NewRequest("POST", "/api/v1/classes", bytes.NewBuffer(testRequestJSON))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	}
}
