package server

import (
	"net/http"
)

func NewServer() http.Handler {
	// Create Server Mux
	mux := http.NewServeMux()

	// Setup Routes
	RegisterRoutes(mux)

	// Add Global Middleware
	middlewareChain := MiddlewareChain(
		RequestLoggerMiddleware,
		RequireAuthMiddleware,
	)

	// Return Handler w/ middleware
	return middlewareChain(mux)
}