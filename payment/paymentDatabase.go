package payment

import "time"

type Payment struct {
	ID             uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	UserID         int       `gorm:"column:user_id;type:varchar(255)"`
	BankID         int       `gorm:"column:bank_id;type:varchar(255)"`
	DeliveryID     int       `gorm:"column:delivery_id;type:varchar(255)"`
	TotalPrice     int       `gorm:"column:total_price;type:varchar(255)"`
	PaymentStatus  string    `gorm:"column:payment_status;type:varchar(255)"`
	DeliveryStatus string    `gorm:"column:delivery_status;type:varchar(255)"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime"`
}