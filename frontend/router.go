package frontend

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Elagoht/bloggo/backend/middleware"
	"github.com/Elagoht/bloggo/utils"
	"github.com/go-chi/chi"
)

func FrontendRouter(router *chi.Mux) {
	frontendRouter := chi.NewRouter()

	frontendRouter.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		render(writer, request, "pages/panel/index.html", map[string]any{
			"Title": "Bloggo - Blog With Go",
		})
	})

	frontendRouter.Get("/auth/login", func(writer http.ResponseWriter, request *http.Request) {
		render(writer, request, "pages/auth/login/index.html", map[string]any{
			"Title": "Bloggo - Login",
		})
	})

	// Serve static files last
	fileServer := http.FileServer(http.Dir("static"))
	frontendRouter.Handle("/*", http.StripPrefix("/static", fileServer))

	router.Mount("/", frontendRouter)
}

func render(
	writer http.ResponseWriter,
	request *http.Request,
	templateName string,
	data any,
) {
	middleware.Handle(func(
		writer http.ResponseWriter,
		request *http.Request,
	) *utils.AppError {
		// Get the base template path
		baseTemplate := filepath.Join("frontend", "templates", "base", "index.html")

		// Get the layout template path based on the template name
		layoutPath := filepath.Join("frontend", "templates", "layouts", "panel", "index.html")

		// Get the page template path
		pageTemplate := filepath.Join("frontend", "templates", templateName)

		// Parse all templates
		templates := []string{
			baseTemplate,
			layoutPath,
			pageTemplate,
		}

		log.Printf("Loading templates: %v", templates)

		tmpl, err := template.ParseFiles(templates...)
		if err != nil {
			log.Printf("Error parsing templates: %v", err)
			return utils.ErrInternalServer
		}

		// Set content type header
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Execute the template
		err = tmpl.ExecuteTemplate(writer, "base", data)
		if err != nil {
			log.Printf("Error executing template: %v", err)
			return utils.ErrInternalServer
		}

		return nil
	})(writer, request)
}
