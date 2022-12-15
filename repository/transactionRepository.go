package repository

import (
	"car-rental/model"
	"database/sql"
	"time"
)

func GetAllTransaction(db *sql.DB) (err error, results []model.Transaction) {
	sql := "SELECT id, date_start, date_finish, total_price, status, car_id, customer_id, created_at, updated_at FROM transactions"
	rows, err := db.Query(sql)
	if err != nil {
		return err, nil
	}
	defer rows.Close()
	for rows.Next() {
		var transaction = model.Transaction{}
		err := rows.Scan(&transaction.ID, &transaction.DateStart, &transaction.DateFinish, &transaction.TotalPrice, &transaction.Status, &transaction.CarId, &transaction.CustomerId, &transaction.CreatedAt, &transaction.UpdatedAt)
		if err != nil {
			return err, nil
		}
		results = append(results, transaction)
	}
	return
}

func InsertTransaction(db *sql.DB, transaction model.PostTransaction) (err error) {
	// sql := "INSERT INTO transactions (date_start, date_finish, total_price, status, car_id, customer_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	// errs := db.QueryRow(sql, transaction.DateStart, transaction.DateFinish, transaction.TotalPrice, "In Progress", transaction.CarId, transaction.CustomerId, time.Now().Local(), time.Now().Local())
	sql := "INSERT INTO transactions (date_start, date_finish, total_price, status, car_id, customer_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	errs := db.QueryRow(sql, transaction.DateStart, transaction.DateFinish, 0, "In Progress", transaction.CarId, transaction.CustomerId, time.Now().Local(), time.Now().Local())
	return errs.Err()
}

func InsertReviewTransaction(db *sql.DB, review model.PostReviewTransaction) (err error) {
	sql := "UPDATE transactions SET rating = $1, review = $2, updated_at = $3 WHERE id = $4"
	errs := db.QueryRow(sql, review.Rating, review.Review, time.Now().Local(), review.TransactionId)
	return errs.Err()
}

// func UpdatePerson(db *sql.DB, person model.Admin) (err error) {
// 	sql := "UPDATE person SET first_name = $1, last_name = $2 WHERE id = $3"
// 	errs := db.QueryRow(sql, person.FirstName, person.LastName, person.ID)
// 	return errs.Err()
// }

func DeleteTransaction(db *sql.DB, transaction model.Transaction) (err error) {
	sql := "DELETE from transactions WHERE id = $1"
	errs := db.QueryRow(sql, transaction.ID)
	return errs.Err()
}
