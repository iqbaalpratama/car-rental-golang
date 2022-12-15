package model

type AuthLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role string `json:"role"`
	ID   int    `json:"id"`
}
