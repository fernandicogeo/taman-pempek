package bank

type BankCreateRequest struct {
	BankCode      string  `json:"bank_code" binding:"required"`
	BankUserCode  *string `json:"bank_user_code"`
	BankName      string  `json:"bank_name" binding:"required"`
	BankNumber    string  `json:"bank_number" binding:"required"`
	BankHolder    string  `json:"bank_holder" binding:"required"`
	BankIsDeleted bool    `json:"bank_is_deleted" binding:"required"`
}
