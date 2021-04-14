package errors

import "errors"

var (
	// ErrorAPIServerNotSet on http.Server missing from an APIServer
	ErrorAPIServerNotSet = errors.New("http.Server missing from APIServer")
)
