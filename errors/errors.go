package errors

import "errors"

var (
	// ErrorAPIServerNotSet on http.Server missing from an APIServer.
	ErrorAPIServerNotSet = errors.New("http.Server missing from APIServer")

	// ErrorDBConnectionFailed on db connection fail.
	ErrorDBConnectionFailed = errors.New("failed to connect to database")

	// ErrorDBNotFound on db not found.
	ErrorDBNotFound = errors.New("database not found")

	// ErrorAPIServerDBMissing on db missing from APIServer.
	ErrorAPIServerDBMissing = errors.New("a *gorm.DB is required by APIServer")

	// ErrorAPIInvalidRequestParameters on invalid request parameters.
	ErrorAPIInvalidRequestParameters = errors.New("invalid request parameters")

	// ErrorAPIInternalError on an internal error on an API endpoint.
	ErrorAPIInternalError = errors.New("something went wrong")

	// ErrorAPIEndBeforeStartDate on end date before start date.
	ErrorAPIEndBeforeStartDate = errors.New("start date should be before end date")

	// ErrorAPIClassEventDoeNotExist on querying a non-existent class event.
	ErrorAPIClassEventDoeNotExist = errors.New("class event does not exist")
)
