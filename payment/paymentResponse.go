package payment

type PaymentResponse struct {
	ID             uint64 `json:"id"`
	UserID         int    `json:"user_id"`
	BankID         int    `json:"bank_id"`
	DeliveryID     int    `json:"delivery_id"`
	TotalPrice     int    `json:"total_price"`
	PaymentStatus  string `json:"payment_status"`
	DeliveryStatus string `json:"delivery_status"`
}
