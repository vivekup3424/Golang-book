package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}
func main() {
	router := gin.Default()
	router.GET("/", getTodos)
	router.Run(":8080")
}
