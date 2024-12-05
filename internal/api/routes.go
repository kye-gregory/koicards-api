package api

import (
	"fmt"
	"net/http"

	h "github.com/kye-gregory/koicards-api/internal/api/handlers"
)

type route struct {
	version string
	prefix string
}

func (r *route) calc(method, endpoint string) string {
	if (r.prefix == "") {return fmt.Sprintf("%s /api/%s/%s", method, r.version, endpoint)}
	return fmt.Sprintf("%s /api/%s/%s/%s", method, r.version, r.prefix, endpoint)
}

func RegisterRoutes(app *App, mux *http.ServeMux) http.Handler {
	// Create Sub-Router
	route := route {
		version: "v1",
		prefix: "",
	}

	// Define Handlers & Middleware
	userHandler := h.NewUserHandler(app.UserService, app.AuthService)
	authMiddleware := h.AuthoriseMiddleware(app.AuthService)

	// "Get All" Routes
	mux.HandleFunc(route.calc("GET", "users"), userHandler.GetUsers)

	// Account Routes
	route.prefix = "account"
	mux.HandleFunc(route.calc("POST", "register"),	userHandler.RegisterUser)
	mux.HandleFunc(route.calc("GET", "verify"),	userHandler.VerifyEmail)
	mux.HandleFunc(route.calc("POST", "login"),	userHandler.Login)
	mux.HandleFunc(route.calc("POST", "logout"), authMiddleware(userHandler.Logout))

	// Add Global Middleware
	return h.ApplyGlobalMiddleware(mux, h.RequestLoggerMiddleware)
}