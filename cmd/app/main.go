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
	"github.com/kye-gregory/koicards-api/internal/debug"
	"github.com/kye-gregory/koicards-api/internal/store"
	"github.com/kye-gregory/koicards-api/internal/store/mock"
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
	var errStack debug.ErrorStack
	
	// Watch System Interrupt & Kill Signals
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
    defer cancel()

	// Initialise App
	userStore := mock.NewUserStore()
	db := store.NewDatabase(userStore)
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
			errMsg := fmt.Errorf("%v", err)
			errStack.Add(errMsg)
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
		errMsg := fmt.Errorf("%s", err)
		errStack.Add(errMsg)
	}

	// Return Any Accumulated Errors
	if len(errStack.Errors) > 0 {return &errStack }
	return nil
}


func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}