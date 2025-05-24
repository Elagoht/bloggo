package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Elagoht/bloggo/backend/models"
	"github.com/Elagoht/bloggo/backend/services"
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

func (controller *BlogController) CreateBlog(
	writer http.ResponseWriter,
	request *http.Request,
) *utils.AppError {
	var blog models.RequestBlog
	bodyErr := json.NewDecoder(request.Body).Decode(&blog)
	if bodyErr != nil {
		return utils.ErrBadRequest
	}

	err := controller.service.CreateBlog(blog)
	if err != nil {
		return err
	}

	writer.WriteHeader(http.StatusCreated)
	return nil
}

func (controller *BlogController) PublishBlog(
	writer http.ResponseWriter,
	request *http.Request,
) *utils.AppError {
	slug := chi.URLParam(request, "slug")

	err := controller.service.PublishBlog(slug)
	if err != nil {
		return err
	}

	return nil
}

func (controller *BlogController) UnpublishBlog(
	writer http.ResponseWriter,
	request *http.Request,
) *utils.AppError {
	slug := chi.URLParam(request, "slug")

	err := controller.service.UnpublishBlog(slug)
	if err != nil {
		return err
	}

	return nil
}

func (controller *BlogController) UpdateBlog(
	writer http.ResponseWriter,
	request *http.Request,
) *utils.AppError {
	slug := chi.URLParam(request, "slug")

	var blog models.RequestBlogPartial
	bodyErr := json.NewDecoder(request.Body).Decode(&blog)
	if bodyErr != nil {
		return utils.ErrBadRequest
	}

	err := controller.service.UpdateBlog(slug, blog)
	if err != nil {
		return err
	}

	return nil
}

func (controller *BlogController) DeleteBlog(
	writer http.ResponseWriter,
	request *http.Request,
) *utils.AppError {
	slug := chi.URLParam(request, "slug")

	err := controller.service.DeleteBlog(slug)
	if err != nil {
		return err
	}

	return nil
}
