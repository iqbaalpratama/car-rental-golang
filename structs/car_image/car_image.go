package car_image

import "database/sql"

type CarImage struct {
	ID        int
	ImageURL  string `json:"image_url"`
	CarId     int    `json: "car_id"`
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
