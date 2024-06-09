package user

type UserCreateRequest struct {
	UserName           string `json:"user_name" binding:"required"`
	Email              string `json:"email" binding:"required,email"`
	Password           string `json:"password" binding:"required"`
	UserWhatsappNumber string `json:"user_whatsapp_number" binding:"required"`
	UserGender         string `json:"user_gender" binding:"required"`
}
