package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "2999"
	}

	router := chi.NewRouter()

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, World!")
	})

	log.Println("Starting server on http://localhost:" + port)
	http.ListenAndServe(":"+port, router)
}
