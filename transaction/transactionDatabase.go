package transaction

import "time"

type Transaction struct {
	ID                      uint64
	TransactionCode         string
	TransactionName         *string
	TransactionWhatsapp     *string
	TransactionEmail        *string
	TransactionAddress      *string
	TransactionStatus       *string
	TransactionTotal        *string
	TransactionPaymentProof *string
	TransactionIsDeleted    bool
	CreatedAt               time.Time
	UpdatedAt               time.Time
}
