package server

import "net/http"

func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", s.healthHandler)
	mux.HandleFunc("/api/v1/manifest/", s.manifestHandler)
	mux.HandleFunc("/api/v1/files/", s.fileHandler)

	return mux
}
