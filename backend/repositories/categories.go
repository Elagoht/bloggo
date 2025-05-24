package repositories

import (
	"database/sql"

	"github.com/Elagoht/bloggo/backend/models"
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

func (repository *CategoryRepository) GetAll() ([]models.ResponseCategoryListItem, *utils.AppError) {
	rows, err := repository.dataBase.Query(SQLGetAllCategories)
	if err != nil {
		return nil, utils.MapDatabaseError(err)
	}
	defer rows.Close()

	categories := []models.ResponseCategoryListItem{}

	for rows.Next() {
		var category models.ResponseCategoryListItem
		err := rows.Scan(
			&category.Name,
			&category.Slug,
		)
		if err != nil {
			return nil, utils.MapDatabaseError(err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (repository *CategoryRepository) GetBySlug(
	slug string,
) (models.ResponseCategory, *utils.AppError) {
	row := repository.dataBase.QueryRow(SQLGetCategoryBySlug, slug)

	var category models.ResponseCategory

	err := row.Scan(
		&category.Name,
		&category.Slug,
		&category.Description,
		&category.Keywords,
		&category.Spot,
	)

	if err != nil {
		return models.ResponseCategory{}, utils.MapDatabaseError(err)
	}

	return category, nil
}

func (repository *CategoryRepository) CreateCategory(
	category models.RequestCategory,
) *utils.AppError {
	statement, err := repository.dataBase.Prepare(
		SQLCreateCategory,
	)
	if err != nil {
		return utils.MapDatabaseError(err)
	}
	defer statement.Close()

	_, err = statement.Exec(
		category.Name,
		utils.GenerateSlug(category.Name),
		category.Description,
		category.Keywords,
		category.Spot,
	)
	if err != nil {
		return utils.MapDatabaseError(err)
	}

	return nil
}

func (repository *CategoryRepository) PatchCategory(
	slug string,
	category models.RequestCategoryPartial,
) *utils.AppError {
	fields := []utils.PatchField{
		{Name: "name", Value: category.Name, Skip: category.Name == ""},
		{Name: "slug", Value: utils.GenerateSlug(category.Name), Skip: category.Name == ""},
		{Name: "description", Value: category.Description, Skip: category.Description == ""},
		{Name: "keywords", Value: category.Keywords, Skip: category.Keywords == ""},
		{Name: "spot", Value: category.Spot, Skip: category.Spot == ""},
		{Name: "updatedAt", Value: "datetime('now')", Skip: false},
	}

	query, args := utils.BuildPatchQuery("categories", "slug", slug, fields)

	statement, err := repository.dataBase.Prepare(query)
	if err != nil {
		return utils.MapDatabaseError(err)
	}
	defer statement.Close()

	result, err := statement.Exec(args...)
	if err != nil {
		return utils.MapDatabaseError(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.MapDatabaseError(err)
	}

	if rowsAffected == 0 {
		return utils.ErrNotFound
	}

	return nil
}

func (repository *CategoryRepository) SoftDeleteCategory(
	slug string,
) *utils.AppError {
	result, err := repository.dataBase.Exec(SQLSoftDeleteCategory, slug)
	if err != nil {
		return utils.MapDatabaseError(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.MapDatabaseError(err)
	}

	if rowsAffected == 0 {
		return utils.MapDatabaseError(err)
	}

	return nil
}
