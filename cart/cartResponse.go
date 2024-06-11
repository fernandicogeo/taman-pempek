package cart

type CartResponse struct {
	ID         uint64 `json:"id"`
	UserID     int    `json:"user_id"`
	ProductID  int    `json:"product_id"`
	PaymentID  int    `json:"payment_id"`
	Quantity   int    `json:"quantity"`
	TotalPrice string `json:"total_price"`
	IsActived  string `json:"isActived"`
}