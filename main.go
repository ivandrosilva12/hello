package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

// helloHandler responds with "Hello, World!"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

// healthHandler responds with the health status of the app
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "healthy"}`))
}

// loggerMiddleware logs incoming requests
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	// Set the port
	port := "8085"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	// Create a new ServeMux (router)
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/health", healthHandler)

	// Wrap the router with the logger middleware
	handler := loggerMiddleware(mux)

	// Start the server
	log.Printf("Server is running on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
