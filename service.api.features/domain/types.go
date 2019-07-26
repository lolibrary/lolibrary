package domain

type Feature struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type ListFeaturesResponse struct {
	Features []*Feature `json:"features"`
}
