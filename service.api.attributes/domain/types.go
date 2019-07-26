package domain

type Attribute struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type ListAttributesResponse struct {
	Attributes []*Attribute `json:"attributes"`
}
