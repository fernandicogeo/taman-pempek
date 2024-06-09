package product

import "time"

type Product struct {
	ID                  uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	ProductCode         string    `gorm:"column:product_code;type:varchar(255);not null;unique"`
	ProductUserCode     *string   `gorm:"column:product_user_code;type:varchar(255)"`
	ProductCategoryCode string    `gorm:"column:product_category_code;type:varchar(255);not null"`
	ProductName         string    `gorm:"column:product_name;type:varchar(255);not null"`
	ProductImage        *string   `gorm:"column:product_image;type:varchar(255)"`
	ProductDescription  *string   `gorm:"column:product_description;type:text"`
	ProductPrice        int       `gorm:"column:product_price;not null"`
	ProductStock        int       `gorm:"column:product_stock;not null"`
	ProductIsDeleted    bool      `gorm:"column:product_is_deleted;type:tinyint(1);not null;default:0"`
	CreatedAt           time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt           time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
