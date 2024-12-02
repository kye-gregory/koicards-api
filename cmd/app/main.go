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
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kye-gregory/koicards-api/internal/api"
	errs "github.com/kye-gregory/koicards-api/internal/errors"
	"github.com/kye-gregory/koicards-api/internal/store"
	storePostgres "github.com/kye-gregory/koicards-api/internal/store/postgres"
	storeRedis "github.com/kye-gregory/koicards-api/internal/store/redis"
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
	var wg sync.WaitGroup
	wg.Add(1)
	
	// Watch System Interrupt & Kill Signals
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
    defer cancel()

	// Initialize PostgreSQL (Primary Database)
	dbHost := os.Getenv("POSTGRES_HOST")
    dbPort := os.Getenv("POSTGRES_PORT")
    dbUser := os.Getenv("POSTGRES_USER")
    dbPassword := os.Getenv("POSTGRES_PASSWORD")
    dbName := os.Getenv("POSTGRES_NAME")
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	dbPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil { errs.Internal(errStack, err); return errStack }
	log.Printf("Connecting to DB at %s:%s as %s\n", dbHost, dbPort, dbUser)
	defer dbPool.Close()

	// Initialize Redis (Session Database)
	redisHost := os.Getenv("REDIS_HOST")
    redisPort := os.Getenv("REDIS_PORT")
	rdb := redis.NewClient(&redis.Options{
		Addr: net.JoinHostPort(redisHost, redisPort),
	})
	log.Printf("Connecting to Redis at %s:%s\n", redisHost, redisPort)
	defer rdb.Close()

	// Initialise App
	userStore := storePostgres.NewUserStore(dbPool)
	sessionStore := storeRedis.NewSessionStore(rdb)
	db := store.NewDatabase(userStore, sessionStore)
	app := api.NewApp(db)

	// Create The Server
	httpServer := &http.Server{
		Addr:    net.JoinHostPort("0.0.0.0", "8080"),
		Handler: api.NewRouter(app),
	}

	// Run Server
	wg.Done()
	wg.Wait()
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
	return errStack.Return()
}


func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}