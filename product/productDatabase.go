package product

import "time"

type Product struct {
	ID                  uint64
	ProductCode         string
	ProductUserCode     *string
	ProductCategoryCode string
	ProductName         string
	ProductImage        *string
	ProductDescription  *string
	ProductPrice        int
	ProductStock        int
	ProductIsDeleted    bool
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
