package main

import (
	"fmt"
	"net/http"
)

func HelloWorld() string {
    return "Hello, World!"
}

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}

func main() {
	// Define the port to listen on
	port := ":8000" // Change this to the desired port number

	// Define a handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, I'm Ankush. Welcome to my first Go-application.")
	})

	// Start the HTTP server
	fmt.Printf("Server listening on port %s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}


