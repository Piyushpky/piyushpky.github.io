package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from your Go server!")
	})

	// Get port from environment, default to 8080
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080" // Default port if not set
	}
	port, _ := strconv.Atoi(portStr)

	fmt.Printf("Server listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
