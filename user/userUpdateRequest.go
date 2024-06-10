package user

type UserUpdateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Whatsapp string `json:"whatsapp"`
	Gender   string `json:"gender"`
	Role     string `json:"role"`
}
