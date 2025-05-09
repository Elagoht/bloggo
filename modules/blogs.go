package modules

import (
	"github.com/Elagoht/bloggo/controllers"
	"github.com/Elagoht/bloggo/middleware"
	"github.com/go-chi/chi"
)

func HandleBlogs(router *chi.Mux) {
	blogController := controllers.NewBlogController()

	blogRouter := chi.NewRouter()

	blogRouter.Get("/", middleware.Handle(blogController.GetAllBlogs))
	blogRouter.Get("/{slug}", middleware.Handle(blogController.GetBlogBySlug))
	blogRouter.Post("/", middleware.Handle(blogController.CreateBlog))
	router.Mount("/api/blogs", blogRouter)
}
