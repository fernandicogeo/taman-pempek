package payment

type PaymentResponse struct {
	ID             uint64 `json:"id"`
	UserID         int    `json:"user_id"`
	DeliveryID     int    `json:"delivery_id"`
	TotalPrice     int    `json:"total_price"`
	Image          string `json:"image"`
	PaymentStatus  string `json:"payment_status"`
	DeliveryStatus string `json:"delivery_status"`
}
