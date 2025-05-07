package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Elagoht/bloggo/services"
	"github.com/Elagoht/bloggo/utils"
	"github.com/go-chi/chi"
)

type CategoryController struct {
	categoryService *services.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		categoryService: services.NewCategoryService(),
	}
}

func (controller *CategoryController) GetAllCategories(writer http.ResponseWriter, request *http.Request) {
	categories, err := controller.categoryService.GetAll()
	if err != nil {
		utils.HandleError(err, writer)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(categories)
}

func (controller *CategoryController) GetCategoryBySlug(writer http.ResponseWriter, request *http.Request) {
	slug := chi.URLParam(request, "slug")

	category, err := controller.categoryService.GetBySlug(slug)
	if err != nil {
		utils.HandleError(err, writer)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(category)
}
