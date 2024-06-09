package transactiondetail

type TransactionDetailCreateRequest struct {
	TransactionDetailCode        string `json:"transaction_detail_code" binding:"required"`
	TransactionCode              string `json:"transaction_code" binding:"required"`
	TransactionDetailProductCode string `json:"transaction_detail_product_code" binding:"required"`
	TransactionDetailPrice       int    `json:"transaction_detail_price" binding:"required,number"`
	TransactionDetailQuantity    int    `json:"transaction_detail_quantity" binding:"required,number"`
	TransactionDetailTotalPrice  int    `json:"transaction_detail_total_price" binding:"required,number"`
	TransactionDetailIsDeleted   bool   `json:"transaction_detail_is_deleted" binding:"required"`
}
