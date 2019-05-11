package model

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Salt         string `json:"-"`
	PasswordHash string `json"-"`
	IsAdmin      bool   `json:"is_admin"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
