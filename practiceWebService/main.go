package main

import (
	"log"
	"net/http"
	"os"
	"web-service/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api:s", log.LstdFlags)
	sm := http.NewServeMux()
	//NOTE: registers a handler
	sm.HandleFunc("/hello", handlers.NewHello(logger).ServeHTTP)
	sm.HandleFunc("/goodbye", handlers.NewGoodbye(logger).ServeHTTP)
	err := http.ListenAndServe("localhost:8080", sm)
	if err != nil {
		log.Fatal(err)
	}

}
