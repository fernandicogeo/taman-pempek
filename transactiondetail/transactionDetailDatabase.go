package transactiondetail

import "time"

type TransactionDetail struct {
	ID                           uint64
	TransactionDetailCode        string
	TransactionCode              string
	TransactionDetailProductCode string
	TransactionDetailPrice       int
	TransactionDetailQuantity    int
	TransactionDetailTotalPrice  int
	TransactionDetailIsDeleted   bool
	CreatedAt                    time.Time
	UpdatedAt                    time.Time
}
