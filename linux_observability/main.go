package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := "My friendo!"
		_, err := w.Write([]byte(msg))
		if err != nil {
			log.Fatal("Error writing the response")
		}
	})

	http.ListenAndServe("localhost:4001", router)
}
