package main

import (
	"log"
	"net/http"
)

func main() {
	//v, _ := os.ReadFile("./public/index.html")
	//fmt.Println((string)(v))
	router := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./public"))
	router.Handle("/", fileServer)
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	log.Println("server is listening on port 8080")
	server.ListenAndServe()
}
