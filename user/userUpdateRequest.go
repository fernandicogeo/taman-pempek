package user

type UserUpdateRequest struct {
	UserName           string `json:"user_name"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	UserWhatsappNumber string `json:"user_whatsapp_number"`
	UserGender         string `json:"user_gender"`
}
