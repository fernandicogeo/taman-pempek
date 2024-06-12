package payment

import "encoding/json"

type PaymentUpdateRequest struct {
	BankID         json.Number `json:"bank_id,omitempty"`
	DeliveryID     json.Number `json:"delivery_id,omitempty"`
	TotalPrice     json.Number `json:"total_price,omitempty"`
	PaymentStatus  string      `json:"payment_status,omitempty"`
	DeliveryStatus string      `json:"delivery_status,omitempty"`
}
