package payment

import "encoding/json"

type PaymentUpdateRequest struct {
	BankID         json.Number `form:"bank_id,omitempty"`
	DeliveryID     json.Number `form:"delivery_id,omitempty"`
	TotalPrice     json.Number `form:"total_price,omitempty"`
	PaymentStatus  string      `form:"payment_status,omitempty"`
	DeliveryStatus string      `form:"delivery_status,omitempty"`
}
