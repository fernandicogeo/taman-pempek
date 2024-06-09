package transactiondetail

import "time"

type TransactionDetail struct {
	ID                           uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	TransactionDetailCode        string    `gorm:"column:transaction_detail_code;type:varchar(255);not null"`
	TransactionCode              string    `gorm:"column:transaction_code;type:varchar(255);not null"`
	TransactionDetailProductCode string    `gorm:"column:transaction_detail_product_code;type:varchar(255);not null"`
	TransactionDetailPrice       int       `gorm:"column:transaction_detail_price;not null"`
	TransactionDetailQuantity    int       `gorm:"column:transaction_detail_quantity;not null"`
	TransactionDetailTotalPrice  int       `gorm:"column:transaction_detail_total_price;not null"`
	TransactionDetailIsDeleted   bool      `gorm:"column:transaction_detail_is_deleted;type:tinyint(1);not null;default:0"`
	CreatedAt                    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt                    time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
