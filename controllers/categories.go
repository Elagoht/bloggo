package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Elagoht/bloggo/models"
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

func (controller *CategoryController) GetAllCategories(
	writer http.ResponseWriter,
	request *http.Request,
) {
	categories, err := controller.categoryService.GetAll()
	if err != nil {
		utils.HandleError(err, writer)
		return
	}

	json.NewEncoder(writer).Encode(categories)
}

func (controller *CategoryController) GetCategoryBySlug(
	writer http.ResponseWriter,
	request *http.Request,
) {
	slug := chi.URLParam(request, "slug")

	category, err := controller.categoryService.GetBySlug(slug)
	if err != nil {
		utils.HandleError(err, writer)
		return
	}

	json.NewEncoder(writer).Encode(category)
}

func (controller *CategoryController) CreateCategory(
	writer http.ResponseWriter,
	request *http.Request,
) {
	var category models.RequestCategoryCreate
	err := json.NewDecoder(request.Body).Decode(&category)
	if err != nil {
		utils.HandleError(err, writer)
		return
	}

	createdCategory, err := controller.categoryService.Create(category)
	if err != nil {
		if err.Error() == "category already exists" {
			utils.NewErrorWithStatus(
				utils.NewError("category already exists"),
				writer,
				http.StatusConflict,
			)
		} else {
			utils.HandleError(err, writer)
		}
		return
	}

	json.NewEncoder(writer).Encode(createdCategory)
}

func (controller *CategoryController) UpdateCategory(
	writer http.ResponseWriter,
	request *http.Request,
) {
	slug := chi.URLParam(request, "slug")

	var category models.RequestCategoryUpdate
	err := json.NewDecoder(request.Body).Decode(&category)
	if err != nil {
		utils.HandleError(err, writer)
		return
	}

	updatedCategory, err := controller.categoryService.Update(slug, category)
	if err != nil {
		utils.HandleError(err, writer)
		return
	}

	json.NewEncoder(writer).Encode(updatedCategory)
}

func (controller *CategoryController) PatchCategory(
	writer http.ResponseWriter,
	request *http.Request,
) {
	slug := chi.URLParam(request, "slug")

	var category models.RequestCategoryUpdate
	err := json.NewDecoder(request.Body).Decode(&category)
	if err != nil {
		utils.HandleError(err, writer)
		return
	}

	patchedCategory, err := controller.categoryService.Patch(slug, category)
	if err != nil {
		utils.HandleError(err, writer)
		return
	}

	json.NewEncoder(writer).Encode(patchedCategory)
}

func (controller *CategoryController) DeleteCategory(
	writer http.ResponseWriter,
	request *http.Request,
) {
	slug := chi.URLParam(request, "slug")

	err := controller.categoryService.Delete(slug)
	if err != nil {
		utils.HandleError(err, writer)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
