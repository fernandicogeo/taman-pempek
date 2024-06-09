package user

type UserCreateRequest struct {
	UserCode           string `json:"user_code" binding:"required"`
	Role               string `json:"role" binding:"required"`
	UserName           string `json:"user_name" binding:"required"`
	Email              string `json:"email" binding:"required,email"`
	Password           string `json:"password" binding:"required"`
	UserWhatsappNumber string `json:"user_whatsapp_number" binding:"required"`
	UserGender         string `json:"user_gender" binding:"required"`
	UserIsDeleted      bool   `json:"user_is_deleted" binding:"required"`
}
