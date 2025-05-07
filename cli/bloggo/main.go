package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Elagoht/bloggo/middleware"
	"github.com/Elagoht/bloggo/modules"
	"github.com/Elagoht/bloggo/utils"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "2999"
	}

	utils.InitDB()

	router := chi.NewRouter()

	// Apply the middleware
	router.Use(middleware.JsonContentTypeMiddleware)

	router = modules.HandleCategories(router)

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, World!")
	})

	log.Println("Starting server on http://localhost:" + port)
	http.ListenAndServe(":"+port, router)
}
