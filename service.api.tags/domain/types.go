package domain

type Tag struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type ListTagsResponse struct {
	Tags []*Tag `json:"tags"`
}
