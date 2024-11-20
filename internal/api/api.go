package api

import (
	"net/http"

	"github.com/kye-gregory/koicards-api/internal/services"
	"github.com/kye-gregory/koicards-api/internal/store"
)

type App struct {
	DB          *store.Database
	UserService *services.UserService
	AuthService *services.AuthService
}

func NewApp(db *store.Database) *App {
	return &App{
		DB:          db,
		UserService: services.NewUserService(db.UserStore),
		AuthService: services.NewAuthService(),
	}
}

func NewRouter(app *App) http.Handler {
	// Create Server Mux
	mux := http.NewServeMux()

	// Setup Routes
	RegisterRoutes(app, mux)

	// Add Global Middleware
	middlewareChain := MiddlewareChain(
		RequestLoggerMiddleware,
	)

	// Return Handler w/ middleware
	return middlewareChain(mux)
}