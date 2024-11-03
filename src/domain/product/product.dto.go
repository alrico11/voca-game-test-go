package product

type CreateProductDTO struct {
	Availability int     `json:"availability"`
	Description  *string `json:"description"`
	Name         string  `json:"name"`
	Price        int     `json:"price"`
}

type UpdateProductDTO struct {
	Availability *int    `json:"availability,omitempty"`
	Description  *string `json:"description,omitempty"`
	Name         *string `json:"name,omitempty"`
	Price        *int    `json:"price,omitempty"`
}

type FindAllProductQuery struct {
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Search string `json:"search"`
	Sort   string `json:"sort"`   // asc or desc
	SortBy string `json:"sortBy"` // field to sort by
}
