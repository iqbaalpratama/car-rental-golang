package admin

import "database/sql"

type Admin struct {
	ID          int
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_no"`
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}
