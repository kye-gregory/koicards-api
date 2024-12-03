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
			// Get Cookie
			httpStack := errpkg.NewHttpStack().WithStatus(http.StatusUnauthorized)
			structuredErr := errs.AuthUnauthorised("invalid session id")
			sessionCookie, err := r.Cookie("session_id")
			if err != nil { httpStack.Add(structuredErr) }
			if returnHttpError(w, httpStack) { return }

			// Check Database
			httpStack.WithStatus(http.StatusUnauthorized)
			auth.VerifySession(sessionCookie.Value, httpStack)
			if returnHttpError(w, httpStack) { return }

			next.ServeHTTP(w,r)
		})
	}
}