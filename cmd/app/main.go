package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kye-gregory/koicards-api/internal/api"
	errs "github.com/kye-gregory/koicards-api/internal/errors"
	"github.com/kye-gregory/koicards-api/internal/store"
	"github.com/kye-gregory/koicards-api/internal/store/mock"
	errpkg "github.com/kye-gregory/koicards-api/pkg/debug/errors"
)


func run(
	ctx    	context.Context,
	args   	[]string,
	getenv 	func(string) string,
	stdin  	io.Reader,
	stdout 	io.Writer,
	stderr 	io.Writer,
) error {
	// Initialise Environment
	time.Local = time.UTC
	log.SetOutput(stderr)
	errStack := errpkg.NewStack()
	
	// Watch System Interrupt & Kill Signals
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
    defer cancel()

	// Initialise App
	userStore := mock.NewUserStore()
	sessionStore := mock.NewSessionStore()
	db := store.NewDatabase(userStore, sessionStore)
	app := api.NewApp(db)

	// Create The Server
	httpServer := &http.Server{
		Addr:    net.JoinHostPort("localhost", "8080"),
		Handler: api.NewRouter(app),
	}

	// Run Server
	go func() {
		log.Println("Listening for requests on", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errs.Internal(errStack, err)
			cancel()
		}
	}()

	// Wait For Server
	<-ctx.Done()

	// Add Shutdown Timeout
	shutdownCtx, _ := context.WithTimeout(context.Background(), 10 * time.Second)

	// Shutdown Server
	log.Println("Server shutting down...")
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		errs.Internal(errStack, err)
	}

	// Return Any Errors
	return nil
}


func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}