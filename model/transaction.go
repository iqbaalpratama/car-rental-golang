package model

import "time"

type Transaction struct {
	ID         int       `json:"id"`
	DateStart  string    `json:"date_start"`
	DateFinish string    `json:"date_finish"`
	TotalPrice int       `json:"total_price"`
	Status     string    `json:"status"`
	Rating     int       `json:"rating"`
	Review     string    `json:"review"`
	CarId      int       `json:"car_id"`
	CustomerId int       `json:"customer_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type PostTransaction struct {
	DateStart  string `json:"date_start"`
	DateFinish string `json:"date_finish"`
	TotalPrice int    `json:"total_price"`
	CarId      int    `json:"car_id"`
	CustomerId int    `json:"customer_id"`
}

type PostReviewTransaction struct {
	Rating int    `json:"rating"`
	Review string `json:"review"`
}
