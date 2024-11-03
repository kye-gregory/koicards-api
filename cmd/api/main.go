package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"time"
)

func run(
	ctx    	context.Context,
	args   	[]string,
	getenv 	func(string) string,
	stdin  	io.Reader,
	stdout 	io.Writer,
	stderr 	io.Writer,
) error {
	// Watch System Interrupt
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
    defer cancel()

	// Mock Execution in GoRoutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Fprintln(stdout, "Working...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Wait For GoRoutine
	<- ctx.Done()

	// Catches System Interrupt
    if ctx.Err() == nil { return nil }
	return errors.New("operation interrupted by system signal")
}


func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}