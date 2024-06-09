package productcategory

type ProductCategoryUpdateRequest struct {
	ProductCategoryCode        string  `json:"product_category_code"`
	ProductCategoryName        string  `json:"product_category_name"`
	ProductCategoryImage       *string `json:"product_category_image"`
	ProductCategoryDescription *string `json:"product_category_description"`
	ProductCategoryIsDeleted   bool    `json:"product_category_is_deleted"`
}