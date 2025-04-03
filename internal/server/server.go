package server

import (
	"context"
	"exchanges/internal/config"
	"exchanges/internal/handlers"
	"net"
	"net/http"
)

type Server struct {
	app    *http.Server
	config *config.Config
}

func New(ctx context.Context, config *config.Config) *Server {

	mux := http.NewServeMux()
	mux.HandleFunc("/exchange", handlers.ExchangeHandler)

	app := &http.Server{
		Addr:    config.Server.Address,
		Handler: mux,
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
	}

	return &Server{
		app:    app,
		config: config,
	}
}

func (s *Server) Start() error {
	if err := s.app.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.app.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
