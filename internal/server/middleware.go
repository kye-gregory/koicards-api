package server

import (
	"log"
	"net/http"
)

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middlewares ... Middleware) Middleware {
	return func (next http.Handler) http.HandlerFunc  {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}

		return next.ServeHTTP
	}
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		log.Printf("Method %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w,r)
	}
}

func RequireAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "ValidAuthToken" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w,r)
	}
}