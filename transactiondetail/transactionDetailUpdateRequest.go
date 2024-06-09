package transactiondetail

type TransactionDetailUpdateRequest struct {
	TransactionDetailCode        string `json:"transaction_detail_code"`
	TransactionCode              string `json:"transaction_code"`
	TransactionDetailProductCode string `json:"transaction_detail_product_code"`
	TransactionDetailPrice       int    `json:"transaction_detail_price"`
	TransactionDetailQuantity    int    `json:"transaction_detail_quantity"`
	TransactionDetailTotalPrice  int    `json:"transaction_detail_total_price"`
	TransactionDetailIsDeleted   bool   `json:"transaction_detail_is_deleted"`
}
