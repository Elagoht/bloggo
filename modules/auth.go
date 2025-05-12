package modules

import (
	"github.com/Elagoht/bloggo/controllers"
	"github.com/Elagoht/bloggo/middleware"
	"github.com/go-chi/chi"
)

func HandleAuth(router *chi.Mux) {
	authController := controllers.NewAuthController()

	authRouter := chi.NewRouter()

	authRouter.Post("/login", middleware.Handle(authController.Login))

	router.Mount("/api/auth", authRouter)
}
