package repositories

import (
	"database/sql"

	"github.com/Elagoht/bloggo/models"
	"github.com/Elagoht/bloggo/utils"
)

type CategoryRepository struct {
	dataBase *sql.DB
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		dataBase: utils.DataBase,
	}
}

func (repository *CategoryRepository) GetAll() ([]models.Category, error) {
	rows, err := repository.dataBase.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []models.Category{}

	for rows.Next() {
		var category models.Category
		err := rows.Scan(
			&category.Id,
			&category.Name,
			&category.Slug,
			&category.Description,
			&category.Keywords,
			&category.Spot,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (repository *CategoryRepository) GetBySlug(slug string) (models.Category, error) {
	row := repository.dataBase.QueryRow("SELECT * FROM categories WHERE slug = ?", slug)

	var category models.Category

	err := row.Scan(
		&category.Id,
		&category.Name,
		&category.Slug,
		&category.Description,
		&category.Keywords,
		&category.Spot,
	)

	if err != nil {
		return models.Category{}, utils.NewError("category not found")
	}

	return category, nil
}
