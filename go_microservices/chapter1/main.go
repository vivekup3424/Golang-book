package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type HelloWorldResponse struct {
	Message string `json:"message"`
}
type HelloWorldRequest struct {
	Name string `json:"name"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloWorldResponse{Message: "Hello, World!"}
	data, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
func PostHelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	req := HelloWorldRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	resp := HelloWorldResponse{Message: "Hello, " + req.Name + "!"}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(&resp)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func helloWorldHandlerEncoder(w http.ResponseWriter, r *http.Request) {
	resp := HelloWorldResponse{Message: "Hello, Encoded World!"}
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&resp)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
func main() {
	port := "8080"
	router := http.NewServeMux()
	router.HandleFunc("/hello", helloWorldHandler)
	router.HandleFunc("/hello/encoder", helloWorldHandlerEncoder)
	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Server start on 127:0:0:1:%v\n", port)
	}
}
