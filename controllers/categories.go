package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Elagoht/bloggo/services"
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
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(categories)
}
