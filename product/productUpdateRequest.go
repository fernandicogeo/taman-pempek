package product

type ProductUpdateRequest struct {
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}
