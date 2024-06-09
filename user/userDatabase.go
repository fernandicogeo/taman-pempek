package user

import "time"

type User struct {
	ID                 uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	UserCode           string    `gorm:"column:user_code;type:varchar(255);not null;unique"`
	Role               string    `gorm:"column:role;type:varchar(255);not null;default:SELLER"`
	UserName           string    `gorm:"column:user_name;type:varchar(255);not null"`
	Email              string    `gorm:"column:email;type:varchar(255);not null;unique"`
	Password           string    `gorm:"column:password;type:varchar(255);not null"`
	UserWhatsappNumber string    `gorm:"column:user_whatsapp_number;type:varchar(255);not null"`
	UserGender         string    `gorm:"column:user_gender;type:varchar(255);not null"`
	UserIsDeleted      bool      `gorm:"column:user_is_deleted;type:tinyint(1);not null;default:0"`
	CreatedAt          time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt          time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
