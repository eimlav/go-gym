package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/eimlav/go-gym/api/router"

	"github.com/eimlav/go-gym/config"
	"github.com/eimlav/go-gym/errors"
)

// Server is responsible for running the API.
type Server struct {
	server *http.Server
}

// NewAPIServer creates a new Server instance.
func NewAPIServer() (*Server, error) {
	cfg := config.Cfg
	router := router.NewRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Server.Address, cfg.Server.Port),
		Handler: router,
	}

	return &Server{
		server: server,
	}, nil
}

// Start runs the Server.
func (s *Server) Start() error {
	if s.server == nil {
		return errors.ErrorAPIServerNotSet
	}

	if err := s.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

// Shutdown shuts down the Server.
func (s *Server) Shutdown() error {
	if s.server == nil {
		return errors.ErrorAPIServerNotSet
	}

	return s.server.Shutdown(context.Background())
}

// GetAddress returns the address of the http.Server.
func (s *Server) GetAddress() string {
	if s.server == nil {
		return ""
	}

	return s.server.Addr
}
