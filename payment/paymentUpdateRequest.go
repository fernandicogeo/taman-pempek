package payment

type PaymentUpdateRequest struct {
	// BankID         json.Number `form:"bank_id,omitempty"`
	DeliveryID     int    `form:"delivery_id,omitempty"`
	TotalPrice     int    `form:"total_price,omitempty"`
	PaymentStatus  string `form:"payment_status,omitempty"`
	DeliveryStatus string `form:"delivery_status,omitempty"`
}
