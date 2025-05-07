package modules

import (
	"github.com/Elagoht/bloggo/controllers"
	"github.com/Elagoht/bloggo/middleware"
	"github.com/go-chi/chi"
)

func HandleCategories(router *chi.Mux) *chi.Mux {
	categoryController := controllers.NewCategoryController()

	router.Get("/api/categories", middleware.Handle(categoryController.GetAllCategories))
	router.Get("/api/categories/{slug}", middleware.Handle(categoryController.GetCategoryBySlug))
	router.Post("/api/categories", middleware.Handle(categoryController.CreateCategory))
	router.Put("/api/categories/{slug}", middleware.Handle(categoryController.UpdateCategory))
	router.Patch("/api/categories/{slug}", middleware.Handle(categoryController.PatchCategory))
	router.Delete("/api/categories/{slug}", middleware.Handle(categoryController.DeleteCategory))

	return router
}
