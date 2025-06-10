package models

type User struct {
	ID           int    `json:"id"` //To convert data from DB to user
	UserName     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
