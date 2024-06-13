package bank

type BankCreateRequest struct {
	Type   string `json:"type" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Number string `json:"number" binding:"required"`
}
