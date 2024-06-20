package payment

import (
	"encoding/json"
	"time"
)

type Payment struct {
	// BankID         json.Number `gorm:"column:bank_id;type:varchar(255)"`
	ID             uint64      `gorm:"column:id;primaryKey;autoIncrement"`
	UserID         int         `gorm:"column:user_id;type:varchar(255)"`
	DeliveryID     json.Number `gorm:"column:delivery_id;type:varchar(255)"`
	TotalPrice     json.Number `gorm:"column:total_price;type:varchar(255)"`
	Image          string      `gorm:"column:image;type:varchar(255)"`
	PaymentStatus  string      `gorm:"column:payment_status;type:varchar(255)"`
	DeliveryStatus string      `gorm:"column:delivery_status;type:varchar(255)"`
	CreatedAt      time.Time   `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time   `gorm:"column:updated_at;autoUpdateTime"`
}
