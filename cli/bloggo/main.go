package main

import (
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
	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the port from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "2999"
	}

	// Initialize the database
	utils.InitDB()

	router := chi.NewRouter()

	// Apply the middleware
	router.Use(middleware.JsonContentTypeMiddleware)

	// Mount the modules
	modules.HandleCategories(router)
	modules.HandleBlogs(router)

	// Start the server
	log.Println("Starting server on http://localhost:" + port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
