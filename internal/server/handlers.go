package server

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func (s *Server) manifestHandler(w http.ResponseWriter, r *http.Request) {
	profile := strings.TrimPrefix(
		r.URL.Path,
		"/api/v1/manifest/",
	)

	if profile == "" {
		http.Error(
			w,
			"profile is required",
			http.StatusBadRequest,
		)
		return
	}

	manifest, err := s.storage.Load(profile)
	if err != nil {
		http.Error(
			w,
			"manifest not found",
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(manifest)
}

func (s *Server) fileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rest := strings.TrimPrefix(r.URL.Path, "/api/v1/files/")
	if rest == "" {
		http.Error(w, "profile and path are required", http.StatusBadRequest)
		return
	}

	profileName, relPath, ok := strings.Cut(rest, "/")
	if !ok || profileName == "" || relPath == "" {
		http.Error(w, "profile and path are required", http.StatusBadRequest)
		return
	}

	profile, err := s.config.Profile(profileName)
	if err != nil {
		http.Error(w, "profile not found", http.StatusNotFound)
		return
	}

	relPath = filepath.FromSlash(relPath)
	fullPath := filepath.Join(profile.Root, relPath)

	// protect from path traversal (../../etc/passwd)
	cleanRoot := filepath.Clean(profile.Root)
	cleanFull := filepath.Clean(fullPath)
	if cleanFull != cleanRoot && !strings.HasPrefix(cleanFull, cleanRoot+string(os.PathSeparator)) {
		http.Error(w, "invalid path", http.StatusBadRequest)
		return
	}

	info, err := os.Stat(cleanFull)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "file not found", http.StatusNotFound)
			return
		}
		http.Error(w, "filed to read file", http.StatusInternalServerError)
		return
	}

	if info.IsDir() {
		http.Error(w, "is a directory", http.StatusBadRequest)
		return
	}

	http.ServeFile(w, r, cleanFull)
}
