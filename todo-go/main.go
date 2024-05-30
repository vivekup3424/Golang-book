package main

import (
	"encoding/json"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Buy groceries", Completed: false},
	{ID: "2", Item: "Read a book", Completed: false},
	{ID: "3", Item: "Write code", Completed: false},
	{ID: "4", Item: "Go for a walk", Completed: false},
	{ID: "5", Item: "Cook dinner", Completed: false},
	{ID: "6", Item: "Clean the house", Completed: false},
	{ID: "7", Item: "Call a friend", Completed: false},
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		val, _ := json.Marshal(todos)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(val)
	})
	http.ListenAndServe("localhost:8081", router)
}
