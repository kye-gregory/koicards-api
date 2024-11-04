package api

import (
	"net/http"

	h "github.com/kye-gregory/koicards-api/internal/api/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /register", 	h.HandleUserRegister)
}