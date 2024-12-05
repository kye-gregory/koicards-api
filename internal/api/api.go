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
		AuthService: services.NewAuthService(db.SessionStore),
	}
}

func NewRouter(app *App) http.Handler {
	mux := http.NewServeMux()
	return RegisterRoutes(app, mux)
}