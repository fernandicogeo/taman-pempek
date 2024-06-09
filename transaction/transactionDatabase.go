package transaction

import "time"

type Transaction struct {
	ID                      uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	TransactionCode         string    `gorm:"column:transaction_code;type:varchar(255);not null;unique"`
	TransactionName         *string   `gorm:"column:transaction_name;type:varchar(255)"`
	TransactionWhatsapp     *string   `gorm:"column:transaction_whatsapp;type:varchar(255)"`
	TransactionEmail        *string   `gorm:"column:transaction_email;type:varchar(255)"`
	TransactionAddress      *string   `gorm:"column:transaction_address;type:text"`
	TransactionStatus       *string   `gorm:"column:transaction_status;type:varchar(255)"`
	TransactionTotal        *string   `gorm:"column:transaction_total;type:varchar(255)"`
	TransactionPaymentProof *string   `gorm:"column:transaction_payment_proof;type:varchar(255)"`
	TransactionIsDeleted    bool      `gorm:"column:transaction_is_deleted;type:tinyint(1);not null;default:0"`
	CreatedAt               time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt               time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
