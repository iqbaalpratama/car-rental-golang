package model

import "time"

type Brand struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostPutBrand struct {
	Name string `json:"name"`
}
