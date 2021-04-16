package api

import (
	"net/http"
	"testing"
	"time"

	"github.com/eimlav/go-gym/config"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	err := config.GetConfig("../")
	assert.NoError(t, err)

	apiServer, err := NewAPIServer()
	assert.NoError(t, err)

	go func() {
		time.Sleep(100 * time.Millisecond)

		assert.Equal(t, apiServer.GetAddress(), "0.0.0.0:8080")

		apiServer.Shutdown()
	}()

	err = apiServer.Start()
	assert.ErrorIs(t, err, http.ErrServerClosed)
}
