package user

type UserResponse struct {
	ID                 uint64 `json:"id"`
	UserCode           string `json:"user_code"`
	Role               string `json:"role"`
	UserName           string `json:"user_name"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	UserWhatsappNumber string `json:"user_whatsapp_number"`
	UserGender         string `json:"user_gender"`
	UserIsDeleted      bool   `json:"user_is_deleted"`
}
