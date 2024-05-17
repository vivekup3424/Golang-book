package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve the HTML content
		html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Simple HTML Page</title>
		</head>
		<body>
			<h1>Hello, World!</h1>
			<p>This is a simple HTML page served by a Go server.</p>
		</body>
		</html>
		`
		fmt.Fprintf(w, html)
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":9999", nil)
}
