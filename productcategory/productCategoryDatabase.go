package productcategory

import "time"

type ProductCategory struct {
	ID                         uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	ProductCategoryCode        string    `gorm:"column:product_category_code;type:varchar(255);not null;unique"`
	ProductCategoryName        string    `gorm:"column:product_category_name;type:varchar(255);not null"`
	ProductCategoryImage       *string   `gorm:"column:product_category_image;type:varchar(255)"`
	ProductCategoryDescription *string   `gorm:"column:product_category_description;type:text"`
	ProductCategoryIsDeleted   bool      `gorm:"column:product_category_is_deleted;type:tinyint(1);not null;default:0"`
	CreatedAt                  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt                  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
