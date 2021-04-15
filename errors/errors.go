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
)
