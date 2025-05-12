package frontend

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Elagoht/bloggo/middleware"
	"github.com/Elagoht/bloggo/utils"
	"github.com/go-chi/chi"
)

func FrontendRouter(router *chi.Mux) {
	frontendRouter := chi.NewRouter()

	frontendRouter.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		render(writer, request, "pages/index.html", map[string]any{
			"Title": "Bloggo - Blog With Go",
		})
	})

	frontendRouter.Get("/auth/login", func(writer http.ResponseWriter, request *http.Request) {
		render(writer, request, "pages/login/index.html", map[string]any{
			"Title": "Bloggo - Login",
		})
	})

	// Serve static files last
	fileServer := http.FileServer(http.Dir("statics"))
	frontendRouter.Handle("/*", http.StripPrefix("/statics", fileServer))

	router.Mount("/", frontendRouter)
}

func render(
	writer http.ResponseWriter,
	request *http.Request,
	templateName string,
	data any,
) {
	middleware.Handle(func(writer http.ResponseWriter, request *http.Request) *utils.AppError {
		// Parse both templates at once
		templates := []string{
			filepath.Join("templates", "base", "index.html"),
			filepath.Join("templates", templateName),
		}

		log.Printf("Loading templates: %v", templates)

		tmpl, err := template.ParseFiles(templates...)
		if err != nil {
			log.Printf("Error parsing templates: %v", err)
			return utils.ErrInternalServer
		}

		// Set content type header
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Execute the base template
		err = tmpl.ExecuteTemplate(writer, "base/index", data)
		if err != nil {
			log.Printf("Error executing template: %v", err)
			return utils.ErrInternalServer
		}

		return nil
	})(writer, request)
}
