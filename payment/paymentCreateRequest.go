package payment

import "encoding/json"

type PaymentCreateRequest struct {
	UserID         int         `json:"user_id"`
	BankID         json.Number `json:"bank_id" binding:"required"`
	DeliveryID     json.Number `json:"delivery_id" binding:"required"`
	TotalPrice     json.Number `json:"total_price" binding:"required"`
	PaymentStatus  string      `json:"payment_status" binding:"required"`
	DeliveryStatus string      `json:"delivery_status" binding:"required"`
}
