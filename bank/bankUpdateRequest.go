package bank

type BankUpdateRequest struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	Number string `json:"number"`
}
