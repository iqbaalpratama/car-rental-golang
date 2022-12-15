package model

import "time"

type Car struct {
	ID        int       `json:"id"`
	CarBrand  string    `json:"car_brand"`
	CarNumber string    `json:"car_number"`
	CarModel  string    `json:"car_model"`
	CarYear   int       `json:"car_year"`
	Status    string    `json:"status"`
	RentPrice int       `json:"rent_price"`
	BrandId   int       `json:"brand_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostPutCar struct {
	CarNumber string `json:"car_number"`
	CarModel  string `json:"car_model"`
	CarYear   int    `json:"car_year"`
	RentPrice int    `json:"rent_price"`
	BrandId   int    `json:"brand_id"`
}
