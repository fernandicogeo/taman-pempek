package productcategory

type ProductCategoryCreateRequest struct {
	ProductCategoryCode        string  `json:"product_category_code" binding:"required"`
	ProductCategoryName        string  `json:"product_category_name" binding:"required"`
	ProductCategoryImage       *string `json:"product_category_image"`
	ProductCategoryDescription *string `json:"product_category_description"`
	ProductCategoryIsDeleted   bool    `json:"product_category_is_deleted" binding:"required"`
}
