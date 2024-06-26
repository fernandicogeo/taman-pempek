package setting

type SettingResponse struct {
	ID          uint64 `json:"id"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Contact     string `json:"contact"`
}
