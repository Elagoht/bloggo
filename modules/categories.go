package modules

import (
	"github.com/Elagoht/bloggo/controllers"
	"github.com/go-chi/chi"
)

func HandleCategories(router *chi.Mux) *chi.Mux {
	categoryController := controllers.NewCategoryController()

	router.Get("/api/categories", categoryController.GetAllCategories)
	router.Get("/api/categories/{slug}", categoryController.GetCategoryBySlug)

	return router
}
