package cart

import "encoding/json"

type CartCreateRequest struct {
	UserID     int         `json:"user_id"`
	ProductID  json.Number `json:"product_id" binding:"required,number"`
	PaymentID  json.Number `json:"payment_id" binding:"required,number"`
	Quantity   json.Number `json:"quantity" binding:"required,number"`
	TotalPrice json.Number `json:"total_price" binding:"required,number"`
	IsActived  string      `json:"isActived" binding:"required"`
}
