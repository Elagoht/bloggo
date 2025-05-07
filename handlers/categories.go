package handlers

import (
	"github.com/Elagoht/bloggo/controllers"
	"github.com/go-chi/chi"
)

var categoryController = controllers.NewCategoryController()

func HandleCategories(router *chi.Mux) {
	router.Get("/categories", categoryController.GetAllCategories)
}
