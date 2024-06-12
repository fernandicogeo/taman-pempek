package cart

type CartUpdateRequest struct {
	ProductID  int    `json:"product_id,omitempty"`
	PaymentID  int    `json:"payment_id,omitempty"`
	Quantity   int    `json:"quantity,omitempty"`
	TotalPrice int    `json:"total_price,omitempty"`
	IsActived  string `json:"isActived,omitempty"`
}
