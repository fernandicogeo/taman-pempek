package payment

import (
	"encoding/json"
	"mime/multipart"
)

type PaymentCreateRequest struct {
	UserID         int                  `json:"user_id"`
	BankID         json.Number          `json:"bank_id" binding:"required"`
	DeliveryID     json.Number          `json:"delivery_id"`
	TotalPrice     json.Number          `json:"total_price" binding:"required"`
	Image          multipart.FileHeader `form:"image" binding:"required"`
	PaymentStatus  string               `json:"payment_status" binding:"required"`
	DeliveryStatus string               `json:"delivery_status" binding:"required"`
}
