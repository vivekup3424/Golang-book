package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"web-service/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api:s", log.LstdFlags)
	router := http.NewServeMux()
	server := http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        router,
		IdleTimeout:    120 * time.Second,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//NOTE: registers a handler
	router.Handle("/hello", handlers.NewHello(logger))
	router.Handle("/goodbye", handlers.NewGoodbye(logger))
	router.Handle("/products", handlers.NewProducts(logger))
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()
	//NOTE: code for managing the shutdown of the server
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	s := <-sigChan
	logger.Println("Received terminate, graceful shutdown", s)
	d := time.Now().Add(30 * time.Second)
	timeoutContext, _ := context.WithDeadline(context.Background(), d)
	err := server.Shutdown(timeoutContext)
	if err != nil {
		logger.Fatal(err)
	}
}
