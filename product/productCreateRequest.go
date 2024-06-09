package product

type ProductCreateRequest struct {
	ProductCode         string  `json:"product_code" binding:"required"`
	ProductUserCode     *string `json:"product_user_code"`
	ProductCategoryCode string  `json:"product_category_code" binding:"required"`
	ProductName         string  `json:"product_name" binding:"required"`
	ProductImage        *string `json:"product_image"`
	ProductDescription  *string `json:"product_description"`
	ProductPrice        int     `json:"product_price" binding:"required,number"`
	ProductStock        int     `json:"product_stock" binding:"required,number"`
	ProductIsDeleted    bool    `json:"product_is_deleted" binding:"required"`
}
