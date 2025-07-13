package main

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/genGuide", HandleGenGuide)
}
