package main

import "net/http"

func main() {
	const addr = "localhost:8080"
	router := http.NewServeMux()
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	filesHandler := http.FileServer(http.Dir("config.Static"))
	router.Handle("/static",
		http.StripPrefix("/static/", filesHandler))
	router.HandleFunc("/", index)
	router.HandleFunc("/err", err)
	router.Handle("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server.ListenAndServe()
}
