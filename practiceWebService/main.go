package main

import (
	"log"
	"net/http"
	"os"
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
	router.HandleFunc("/hello", handlers.NewHello(logger).ServeHTTP)
	router.HandleFunc("/goodbye", handlers.NewGoodbye(logger).ServeHTTP)
	err := server.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
