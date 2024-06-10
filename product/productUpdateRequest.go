package product

type ProductUpdateRequest struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}
