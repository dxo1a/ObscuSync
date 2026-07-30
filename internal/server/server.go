package server

import (
	"fmt"
	"net/http"

	"github.com/dxo1a/obscusync/internal/config"
	"github.com/dxo1a/obscusync/internal/manifest"
)

type Server struct {
	http    *http.Server
	storage *manifest.Storage
	config  *config.Manager
}

func New(address string, storage *manifest.Storage, cfg *config.Manager) *Server {
	s := &Server{
		storage: storage,
		config:  cfg,
	}

	s.http = &http.Server{
		Addr:    address,
		Handler: s.routes(),
	}

	return s
}

func (s *Server) Start() error {
	fmt.Printf("ObscuSync server listening on %s\n", s.http.Addr)

	return s.http.ListenAndServe()
}
