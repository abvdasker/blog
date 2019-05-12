package model

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}

func (r *CreateUserRequest) ToUser() *User {
	return NewUser(
		r.Username,
		r.Password,
		r.IsAdmin,
	)
}
