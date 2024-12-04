package handlers

import (
	"log"
	"net/http"

	errs "github.com/kye-gregory/koicards-api/internal/errors"
	"github.com/kye-gregory/koicards-api/internal/services"
	errpkg "github.com/kye-gregory/koicards-api/pkg/debug/errors"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc


// Predefine Chains
var GlobalMiddlewareChain = MiddlewareChain(RequestLoggerMiddleware)


func ApplyGlobalMiddleware(mux *http.ServeMux, chain Middleware) http.HandlerFunc {
	muxWithMiddleware := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r)
	})

	return chain(muxWithMiddleware)
}


func MiddlewareChain(middlewares ... Middleware) Middleware {
	return func (finalHandler http.HandlerFunc) http.HandlerFunc  {
		for i := len(middlewares) - 1; i >= 0; i-- {
			finalHandler = middlewares[i](finalHandler)
		}

		return finalHandler
	}
}

func RequestLoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: %s, Path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w,r)
	})
}

func AuthoriseMiddleware(auth *services.AuthService) func(http.HandlerFunc) http.HandlerFunc {
    return func(next http.HandlerFunc) http.HandlerFunc {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Define Error Handling
			httpStack := errpkg.NewHttpStack().WithStatus(http.StatusUnauthorized)
			
			// Get Session Cookie
			structuredErr := errs.LogoutAlreadyLoggedOut("You are already logged out.")
			sessionCookie, err := r.Cookie("session_id")
			if err != nil { httpStack.Add(structuredErr) }

			// Get CSRF Token Cookie
			csrfToken := r.Header.Get("X-CSRF-Token")
			
			// Return Errors
			if returnHttpError(w, httpStack) { return }

			// Check Database
			httpStack.WithStatus(http.StatusUnauthorized)
			auth.VerifySession(sessionCookie.Value, csrfToken, httpStack)
			if returnHttpError(w, httpStack) { return }

			next.ServeHTTP(w,r)
		})
	}
}