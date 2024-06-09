package bank

import "time"

type Bank struct {
	ID            uint64     `gorm:"column:id;primaryKey;autoIncrement"`
	BankCode      string     `gorm:"column:bank_code;type:varchar(255);not null;unique"`
	BankUserCode  *string    `gorm:"column:bank_user_code;type:varchar(255)"`
	BankName      string     `gorm:"column:bank_name;type:varchar(255);not null"`
	BankNumber    string     `gorm:"column:bank_number;type:varchar(255);not null"`
	BankHolder    string     `gorm:"column:bank_holder;type:varchar(255);not null"`
	BankIsDeleted bool       `gorm:"column:bank_is_deleted;type:tinyint(1);not null;default:0"`
	CreatedAt     *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
