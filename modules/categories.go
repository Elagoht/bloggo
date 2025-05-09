package modules

import (
	"github.com/Elagoht/bloggo/controllers"
	"github.com/Elagoht/bloggo/middleware"
	"github.com/go-chi/chi"
)

func HandleCategories(router *chi.Mux) {
	categoryController := controllers.NewCategoryController()

	categoryRouter := chi.NewRouter()

	categoryRouter.Get("/", middleware.Handle(categoryController.GetAllCategories))
	categoryRouter.Get("/{slug}", middleware.Handle(categoryController.GetCategoryBySlug))
	categoryRouter.Post("/", middleware.Handle(categoryController.CreateCategory))
	categoryRouter.Patch("/{slug}", middleware.Handle(categoryController.PatchCategory))
	categoryRouter.Delete("/{slug}", middleware.Handle(categoryController.DeleteCategory))

	router.Mount("/api/categories", categoryRouter)
}
