package payment

type PaymentUpdateRequest struct {
	BankID         int    `json:"bank_id,omitempty"`
	DeliveryID     int    `json:"delivery_id,omitempty"`
	TotalPrice     int    `json:"total_price,omitempty"`
	PaymentStatus  string `json:"payment_status,omitempty"`
	DeliveryStatus string `json:"delivery_status,omitempty"`
}
