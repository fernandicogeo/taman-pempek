package payment

import (
	"encoding/json"
	"mime/multipart"
)

type PaymentCreateRequest struct {
	UserID int `form:"user_id" binding:"required"`
	// BankID         json.Number          `form:"bank_id" binding:"required"`
	DeliveryID     json.Number          `form:"delivery_id"`
	TotalPrice     json.Number          `form:"total_price" binding:"required"`
	Image          multipart.FileHeader `form:"image" binding:"required"`
	PaymentStatus  string               `form:"payment_status" binding:"required"`
	DeliveryStatus string               `form:"delivery_status" binding:"required"`
}
