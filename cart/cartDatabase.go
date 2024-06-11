package cart

import "time"

type Cart struct {
	ID         uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	UserID     int       `gorm:"column:user_id;type:varchar(255)"`
	ProductID  int       `gorm:"column:product_id;type:varchar(255)"`
	PaymentID  int       `gorm:"column:payment_id;type:varchar(255)"`
	Quantity   int       `gorm:"column:quantity;type:varchar(255)"`
	TotalPrice int       `gorm:"column:total_price;type:varchar(255)"`
	IsActived  string    `gorm:"column:isActived;type:varchar(255)"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
