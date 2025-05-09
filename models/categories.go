package models

import "time"

type Category struct {
	// Identifiers
	Id   int    `json:"id" validate:"required"`
	Slug string `json:"slug" validate:"required,max=100,slug"`
	// Name and metadata
	Name        string `json:"name" validate:"required,max=100"`
	Description string `json:"description" validate:"required,max=155"`
	Spot        string `json:"spot" validate:"required,max=75"`
	Keywords    string `json:"keywords" validate:"required,max=255"`
	// Timestamps
	CreatedAt time.Time  `json:"createdAt" validate:"required"`
	UpdatedAt time.Time  `json:"updatedAt" validate:"required"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

type CategoryWithBlogs struct {
	Category
	Blogs []Blog `json:"blogs"`
}

type RequestCategory struct {
	Name        string `json:"name" validate:"required,max=100"`
	Description string `json:"description" validate:"required,max=155"`
	Keywords    string `json:"keywords" validate:"required,max=255"`
	Spot        string `json:"spot" validate:"required,max=75"`
}

type RequestCategoryPartial struct {
	Name        string `json:"name,omitempty" validate:"max=100"`
	Description string `json:"description,omitempty" validate:"max=155"`
	Keywords    string `json:"keywords,omitempty" validate:"max=255"`
	Spot        string `json:"spot,omitempty" validate:"max=75"`
}

type ResponseCategoryListItem struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ResponseCategory struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	Spot        string `json:"spot"`
}
