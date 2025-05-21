package dto

type TelegramInitDataUser struct {
	ID              uint   `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Username        string `json:"username"`
	LanguageCode    string `json:"language_code"`
	AllowsWriteToPm bool   `json:"allows_write_to_pm"`
	PhotoUrl        string `json:"photo_url"`
}
