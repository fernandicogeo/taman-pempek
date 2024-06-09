package product

import "time"

type ProductResponse struct {
	ID                  uint64    `json:"id"`
	ProductCode         string    `json:"product_code"`
	ProductUserCode     *string   `json:"product_user_code"`
	ProductCategoryCode string    `json:"product_category_code"`
	ProductName         string    `json:"product_name"`
	ProductImage        *string   `json:"product_image"`
	ProductDescription  *string   `json:"product_description"`
	ProductPrice        int       `json:"product_price"`
	ProductStock        int       `json:"product_stock"`
	ProductIsDeleted    bool      `json:"product_is_deleted"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
