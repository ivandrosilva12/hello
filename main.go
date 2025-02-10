package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
