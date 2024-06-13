package bank

type BankResponse struct {
	ID     uint64 `json:"id"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	Number string `json:"number"`
}
