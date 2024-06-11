package cart

type CartUpdateRequest struct {
	ProductID  int    `json:"product_id"`
	PaymentID  int    `json:"payment_id"`
	Quantity   int    `json:"quantity"`
	TotalPrice string `json:"total_price"`
	IsActived  string `json:"isActived"`
}
