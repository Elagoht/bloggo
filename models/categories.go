package models

import "time"

type Category struct {
	Id int `json:"id"`

	Name string `json:"name"`
	Slug string `json:"slug"`

	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	Spot        string `json:"spot"`

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

type CategoryWithBlogs struct {
	Category
	Blogs []Blog `json:"blogs"`
}

type RequestCategoryCreate struct {
	Name string `json:"name" validate:"required,max=100"`
	Slug string `json:"slug" validate:"required,max=100,alphanum"`

	Description string `json:"description" validate:"required,max=155"`
	Keywords    string `json:"keywords" validate:"required,max=255"`
	Spot        string `json:"spot" validate:"required,max=75"`
}

type RequestCategoryUpdate struct {
	Name        string `json:"name" validate:"max=100"`
	Slug        string `json:"slug" validate:"max=100,alphanum"`
	Description string `json:"description" validate:"max=155"`
	Keywords    string `json:"keywords" validate:"max=255"`
	Spot        string `json:"spot" validate:"max=75"`
}

type ResponseCategoryListItem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
