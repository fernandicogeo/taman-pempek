package transaction

type TransactionCreateRequest struct {
	TransactionCode         string  `json:"transaction_code" binding:"required"`
	TransactionName         *string `json:"transaction_name"`
	TransactionWhatsapp     *string `json:"transaction_whatsapp"`
	TransactionEmail        *string `json:"transaction_email"`
	TransactionAddress      *string `json:"transaction_address"`
	TransactionStatus       *string `json:"transaction_status"`
	TransactionTotal        *string `json:"transaction_total"`
	TransactionPaymentProof *string `json:"transaction_payment_proof"`
	TransactionIsDeleted    bool    `json:"transaction_is_deleted" binding:"required"`
}
