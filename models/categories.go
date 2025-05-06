package models

type Category struct {
	Id int `json:"id"`

	Name string `json:"name"`
	Slug string `json:"slug"`

	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	Spot        string `json:"spot"`
}

type CategoryWithBlogs struct {
	Category
	Blogs []Blog `json:"blogs"`
}
