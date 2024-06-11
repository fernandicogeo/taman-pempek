package cart

type CartCreateRequest struct {
	UserID     int    `json:"user_id" binding:"required"`
	ProductID  int    `json:"product_id" binding:"required"`
	PaymentID  int    `json:"payment_id" binding:"required"`
	Quantity   int    `json:"quantity" binding:"required"`
	TotalPrice string `json:"total_price" binding:"required"`
	IsActived  string `json:"isActived" binding:"required"`
}
