package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ismoilroziboyev/bookshelf/internal/config"
	"github.com/ismoilroziboyev/bookshelf/internal/services"
	"github.com/ismoilroziboyev/bookshelf/internal/transport/rest/handlers"
)

type HttpServer struct {
	*http.Server
	cfg *config.Config
}

func New(cfg *config.Config, service *services.Service) *HttpServer {
	// initialize handlers
	handlers := handlers.New(cfg, service)

	return &HttpServer{
		Server: &http.Server{
			Addr:           fmt.Sprintf("%s:%s", cfg.HttpHost, cfg.HttpPort),
			ReadTimeout:    time.Second * 15,
			WriteTimeout:   time.Second * 15,
			MaxHeaderBytes: 1 << 20,
			Handler:        handlers.API(),
		},
		cfg: cfg,
	}
}

func (s *HttpServer) Run() error {
	return s.ListenAndServe()
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
