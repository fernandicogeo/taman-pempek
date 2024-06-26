package setting

import "mime/multipart"

type SettingUpdateRequest struct {
	Image       *multipart.FileHeader `form:"image,omitempty"`
	Description string                `form:"description,omitempty"`
	Contact     string                `form:"contact,omitempty"`
}
