package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
	"web-service/handlers"
)

const version = "1.0.0"

// For now, the only configuration settings will be the network port that we want the
// server to listen on, and the name of the current operating environment for the
// application (development, staging, production, etc.). We will read in these
// configuration settings from command-line flags when the application starts.
type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}
type Application struct {
	config  config
	eLogger *log.Logger
	iLogger *log.Logger
}

func main() {
	//declare an instance of config
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	//NOTE: loggers for writing messages to stdout
	iLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	eLogger := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//NOTE: creates a new server
	router := http.NewServeMux()
	server := http.Server{
		Addr:           "127.0.0.1:" + strconv.Itoa(cfg.port),
		Handler:        router,
		IdleTimeout:    120 * time.Second,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//NOTE: registers a handler
	router.Handle("/hello", handlers.NewHello(iLogger))
	router.Handle("/goodbye", handlers.NewGoodbye(iLogger))
	router.Handle("/products", handlers.NewProducts(iLogger))
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			iLogger.Fatal(err)
		}
	}()
	//NOTE: code for managing the shutdown of the server
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	s := <-sigChan
	iLogger.Println("Received terminate, graceful shutdown", s)
	d := time.Now().Add(30 * time.Second)
	timeoutContext, _ := context.WithDeadline(context.Background(), d)
	err := server.Shutdown(timeoutContext)
	if err != nil {
		eLogger.Fatal(err)
	}
}
