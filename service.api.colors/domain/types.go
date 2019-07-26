package domain

type Color struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type ListColorsResponse struct {
	Colors []*Color `json:"colors"`
}
