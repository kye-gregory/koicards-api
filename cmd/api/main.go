package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/kye-gregory/koicards-api/internal/server"
)

func run(
	ctx    	context.Context,
	args   	[]string,
	getenv 	func(string) string,
	stdin  	io.Reader,
	stdout 	io.Writer,
	stderr 	io.Writer,
) error {
	log.SetOutput(stderr)
	
	// Watch System Interrupt
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
    defer cancel()

	// Create The Server
	handler := server.NewServer()
	httpServer := &http.Server{
		Addr:    net.JoinHostPort("localhost", "8080"),
		Handler: handler,
	}

	// Run Server
	go func() {
		log.Println("Listening for requests on", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("ListenAndServe error: %v", err)
			cancel()
		}
	}()

	// Finalize Server
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		// Catch System Interrupt
		<-ctx.Done()

		// Add Shutdown Timeout
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10 * time.Second)
		defer cancel()

		// Shutdown Server
		log.Println("Server shutting down...")
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Printf("error shutting down http server: %s", err)
		}
	}()

	// Exit Run
	wg.Wait()
	return nil
}


func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}