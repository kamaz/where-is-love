package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/kamaz/where-is-love/types"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rs/zerolog"
)

// Server is the main API server
type Server struct {
	echo      *echo.Echo
	port      uint
	logger    zerolog.Logger
	endpoints []types.Endpoint
}

func CreateServer(port uint, endpoints ...types.Endpoint) *Server {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	e := echo.New()
	e.HideBanner = true
	return &Server{
		echo:      e,
		port:      port,
		logger:    logger,
		endpoints: endpoints,
	}
}

// Run starts the server
func (s *Server) Run() {
	log := s.logger
	e := s.echo

	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())

	for _, endpoint := range s.endpoints {
		switch endpoint.Method() {
		case "POST":
			e.POST(endpoint.Path(), endpoint.Process, endpoint.Middlewares()...)
		case "GET":
			e.GET(endpoint.Path(), endpoint.Process, endpoint.Middlewares()...)
		}
	}

	go func() {
		port := fmt.Sprintf(":%d", s.port)

		if e := log.Info(); e.Enabled() {
			e.Str("port", port).Msg("starting server")
		}
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			if e := log.Error(); e.Enabled() {
				e.Err(err).Msg("failed to start server")
			}
		}
	}()
}

// Stop stops the server
func (s *Server) Stop(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}
