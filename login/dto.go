package main

type HandlerResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type LoginRequest struct {
	Username string `json:"username,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

type User struct {
	Username string `json:"-"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RealName string `json:"realname"`
}
