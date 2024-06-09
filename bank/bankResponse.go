package bank

type BankResponse struct {
	ID            uint64  `json:"id"`
	BankCode      string  `json:"bank_code"`
	BankUserCode  *string `json:"bank_user_code"`
	BankName      string  `json:"bank_name"`
	BankNumber    string  `json:"bank_number"`
	BankHolder    string  `json:"bank_holder"`
	BankIsDeleted bool    `json:"bank_is_deleted"`
}
