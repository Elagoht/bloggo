package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Elagoht/bloggo/services"
	"github.com/Elagoht/bloggo/utils"
	"github.com/go-chi/chi"
)

type BlogController struct {
	service *services.BlogService
}

func NewBlogController() *BlogController {
	return &BlogController{
		service: services.NewBlogService(),
	}
}

func (controller *BlogController) GetAllBlogs(
	writer http.ResponseWriter,
	request *http.Request,
) *utils.AppError {
	blogs, err := controller.service.GetAll()
	if err != nil {
		return err
	}

	json.NewEncoder(writer).Encode(blogs)
	return nil
}

func (controller *BlogController) GetBlogBySlug(
	writer http.ResponseWriter,
	request *http.Request,
) *utils.AppError {
	slug := chi.URLParam(request, "slug")

	blog, err := controller.service.GetBySlug(slug)
	if err != nil {
		return err
	}

	json.NewEncoder(writer).Encode(blog)
	return nil
}
