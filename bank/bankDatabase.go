package bank

import "time"

type Bank struct {
	ID            uint64
	BankCode      string
	BankUserCode  *string
	BankName      string
	BankNumber    string
	BankHolder    string
	BankIsDeleted bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
