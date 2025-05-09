package services

import (
	"net/http"

	"github.com/Elagoht/bloggo/models"
	"github.com/Elagoht/bloggo/pipes"
	"github.com/Elagoht/bloggo/repositories"
	"github.com/Elagoht/bloggo/utils"
	"github.com/go-playground/validator/v10"
)

type BlogService struct {
	repository *repositories.BlogRepository
	validate   *validator.Validate
}

func NewBlogService() *BlogService {
	return &BlogService{
		repository: repositories.NewBlogRepository(),
		validate:   pipes.GetValidator(),
	}
}

func (service *BlogService) GetAll() ([]models.ResponseBlogCard, *utils.AppError) {
	return service.repository.GetAll()
}

func (service *BlogService) GetBySlug(slug string) (models.ResponseBlog, *utils.AppError) {
	return service.repository.GetBySlug(slug)
}

func (service *BlogService) CreateBlog(blog models.RequestBlog) *utils.AppError {
	err := service.validate.Struct(blog)
	if err != nil {
		return utils.NewAppError(
			http.StatusBadRequest,
			"Invalid request body",
			err,
		)
	}
	return service.repository.CreateBlog(blog)
}
