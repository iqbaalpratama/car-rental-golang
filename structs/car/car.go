package car

import "database/sql"

type Car struct {
	ID        int
	CarNumber string `json:"car_number"`
	CarModel  string `json:"car_model"`
	CarYear   int    `json:"car_year"`
	RentPrice int    `json:"rent_price"`
	BrandId   int    `json: "brand_id"`
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
