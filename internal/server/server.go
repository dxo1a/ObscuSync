package server

import (
	"fmt"
	"net/http"
)

func Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "OK")
	})

	fmt.Println("GameSync server started")
	fmt.Println("Listening on :8080")

	return http.ListenAndServe(":8080", mux)
}
