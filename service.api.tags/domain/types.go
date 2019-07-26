package domain

type Tag struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type ListTagsResponse struct {
	Tags []*Tag `json:"tags"`
}
