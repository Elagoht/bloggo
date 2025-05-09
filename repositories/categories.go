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

func (repository *CategoryRepository) GetAll() ([]models.ResponseCategoryListItem, *utils.AppError) {
	rows, err := repository.dataBase.Query("SELECT id, name, slug FROM categories WHERE deletedAt IS NULL ORDER BY createdAt DESC")
	if err != nil {
		return nil, utils.MapDatabaseError(err)
	}
	defer rows.Close()

	categories := []models.ResponseCategoryListItem{}

	for rows.Next() {
		var category models.ResponseCategoryListItem
		err := rows.Scan(
			&category.Id,
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

func (repository *CategoryRepository) GetBySlug(slug string) (models.Category, *utils.AppError) {
	row := repository.dataBase.QueryRow("SELECT * FROM categories WHERE deletedAt IS NULL AND slug = ?", slug)

	var category models.Category

	err := row.Scan(
		&category.Id,
		&category.Name,
		&category.Slug,
		&category.Description,
		&category.Keywords,
		&category.Spot,
		&category.CreatedAt,
		&category.UpdatedAt,
		&category.DeletedAt,
	)

	if err != nil {
		return models.Category{}, utils.MapDatabaseError(err)
	}

	return category, nil
}

func (repository *CategoryRepository) Create(
	category models.RequestCategory,
) (models.RequestCategory, *utils.AppError) {
	statement, err := repository.dataBase.Prepare(
		"INSERT INTO categories (name, slug, description, keywords, spot) VALUES (?, ?, ?, ?, ?)",
	)
	if err != nil {
		return models.RequestCategory{}, utils.MapDatabaseError(err)
	}
	defer statement.Close()

	_, err = statement.Exec(
		category.Name,
		category.Slug,
		category.Description,
		category.Keywords,
		category.Spot,
	)
	if err != nil {
		return models.RequestCategory{}, utils.MapDatabaseError(err)
	}

	return category, nil
}

func (repository *CategoryRepository) Update(
	slug string,
	category models.RequestCategory,
) *utils.AppError {
	statement, err := repository.dataBase.Prepare(
		"UPDATE categories SET name = ?, slug = ?, description = ?, keywords = ?, spot = ?, updatedAt = datetime('now') WHERE slug = ?",
	)
	if err != nil {
		return utils.MapDatabaseError(err)
	}
	defer statement.Close()

	_, err = statement.Exec(
		category.Name,
		category.Slug,
		category.Description,
		category.Keywords,
		category.Spot,
		slug,
	)
	if err != nil {
		return utils.MapDatabaseError(err)
	}

	return nil
}

func (repository *CategoryRepository) Patch(
	slug string,
	category models.RequestCategoryPartial,
) *utils.AppError {
	fields := []utils.PatchField{
		{Name: "name", Value: category.Name, Skip: category.Name == ""},
		{Name: "slug", Value: category.Slug, Skip: category.Slug == ""},
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

func (repository *CategoryRepository) Delete(
	slug string,
) *utils.AppError {
	result, err := repository.dataBase.Exec(
		"UPDATE categories SET deletedAt = datetime('now') WHERE slug = ?",
		slug,
	)
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
