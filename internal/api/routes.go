package api

import (
	"net/http"

	h "github.com/kye-gregory/koicards-api/internal/api/handlers"
)

func RegisterRoutes(app *App, mux *http.ServeMux) {
	userHandler := h.NewUserHandler(app.UserService)
	mux.HandleFunc("POST /register", userHandler.RegisterUser)
}