package domain

type Category struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type ListCategoriesResponse struct {
	Categories []*Category `json:"categories"`
}
