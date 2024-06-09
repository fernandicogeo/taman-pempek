package user

import "time"

type User struct {
	ID                 uint64
	UserCode           string
	Role               string
	UserName           string
	Email              string
	Password           string
	UserWhatsappNumber string
	UserGender         string
	UserIsDeleted      bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
