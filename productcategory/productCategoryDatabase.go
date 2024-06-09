package productcategory

import "time"

type ProductCategory struct {
	ID                         uint64
	ProductCategoryCode        string
	ProductCategoryName        string
	ProductCategoryImage       *string
	ProductCategoryDescription *string
	ProductCategoryIsDeleted   bool
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
}
