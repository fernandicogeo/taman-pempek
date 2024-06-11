package product

type ProductCreateRequest struct {
	UserID      int    `json:"user_id" binding:"required"`
	CategoryID  int    `json:"category_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Price       int    `json:"price" binding:"required,number"`
	Stock       int    `json:"stock" binding:"required,number"`
}
