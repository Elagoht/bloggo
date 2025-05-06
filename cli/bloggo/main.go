package main

import (
	"fmt"
	"log"
	"os"

	"net/http"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "2999"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	log.Println("Starting server on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
