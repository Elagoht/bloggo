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
	Category
}

type RequestCategoryUpdate struct {
	Category
}

type ResponseCategoryListItem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
