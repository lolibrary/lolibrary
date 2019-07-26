package domain

type Brand struct {
	Slug      string `json:"slug"`
	ShortName string `json:"short_name"`
	Name      string `json:"name"`
	// Image struct { URL string } should be here
}

type ListBrandsResponse struct {
	Brands []*Brand `json:"brands"`
}
