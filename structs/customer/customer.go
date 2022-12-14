package customer

import "database/sql"

type Customer struct {
	ID          int
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_no"`
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}
