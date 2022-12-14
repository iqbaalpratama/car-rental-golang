package transaction

import "database/sql"

type Transaction struct {
	ID         int
	DateStart  string `json:"date_start"`
	DateFinish string `json:"date_finish"`
	TotalPrice int    `json:"total_price"`
	Status     string `json:"status"`
	CarId      int    `json: "car_id"`
	CustomerId int    `json: "customer_id"`
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
}
