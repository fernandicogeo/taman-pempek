package payment

type PaymentCreateRequest struct {
	UserID         int    `json:"user_id" binding:"required"`
	BankID         int    `json:"bank_id" binding:"required"`
	DeliveryID     int    `json:"delivery_id" binding:"required"`
	TotalPrice     int    `json:"total_price" binding:"required"`
	PaymentStatus  string `json:"payment_status" binding:"required"`
	DeliveryStatus string `json:"delivery_status" binding:"required"`
}
