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
	blogRouter.Patch("/{slug}/publish", middleware.Handle(blogController.PublishBlog))
	blogRouter.Patch("/{slug}/unpublish", middleware.Handle(blogController.UnpublishBlog))
	blogRouter.Patch("/{slug}", middleware.Handle(blogController.UpdateBlog))
	blogRouter.Delete("/{slug}", middleware.Handle(blogController.DeleteBlog))

	router.Mount("/api/blogs", blogRouter)
}
