package review

import "database/sql"

type Review struct {
	ID            int
	Rating        int    `json:"rating"`
	Review        string `json:"review"`
	TransactionId int    `json: "transaction_id"`
	CreatedAt     sql.NullTime
	UpdatedAt     sql.NullTime
}
