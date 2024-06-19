package payment

import "encoding/json"

type PaymentResponse struct {
	ID             uint64      `json:"id"`
	UserID         int         `json:"user_id"`
	BankID         json.Number `json:"bank_id"`
	DeliveryID     json.Number `json:"delivery_id"`
	TotalPrice     json.Number `json:"total_price"`
	Image          string      `json:"image"`
	PaymentStatus  string      `json:"payment_status"`
	DeliveryStatus string      `json:"delivery_status"`
}
