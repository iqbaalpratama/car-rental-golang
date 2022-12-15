package model

import "time"

type CarImage struct {
	ID        int       `json:"id"`
	ImageURL  string    `json:"image_url"`
	CarId     int       `json:"car_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostCarImage struct {
	ImageURL string `json:"image_url"`
	CarId    int    `json:"car_id"`
}
