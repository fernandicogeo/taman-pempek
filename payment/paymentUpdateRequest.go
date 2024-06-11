package payment

type PaymentUpdateRequest struct {
	BankID         int    `json:"bank_id"`
	DeliveryID     int    `json:"delivery_id"`
	TotalPrice     int    `json:"total_price"`
	PaymentStatus  string `json:"payment_status"`
	DeliveryStatus string `json:"delivery_status"`
}
