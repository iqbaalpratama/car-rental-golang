package repository

import (
	"car-rental/model"
	"database/sql"
	"errors"
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

func GetTransactionByCarId(db *sql.DB, carId int) (err error, results []model.Transaction) {
	sql := "SELECT id, date_start, date_finish, total_price, status, car_id, customer_id, created_at, updated_at FROM transactions WHERE car_id = $1"
	rows, err := db.Query(sql, carId)
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

func GetTransactionByCustomerId(db *sql.DB, customerId int) (err error, results []model.Transaction) {
	sql := "SELECT id, date_start, date_finish, total_price, status, car_id, customer_id, created_at, updated_at FROM transactions WHERE customer_id = $1"
	rows, err := db.Query(sql, customerId)
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
	sql := "INSERT INTO transactions (date_start, date_finish, total_price, status, car_id, customer_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	errs := db.QueryRow(sql, transaction.DateStart, transaction.DateFinish, transaction.TotalPrice, "Created", transaction.CarId, transaction.CustomerId, time.Now().Local(), time.Now().Local())
	return errs.Err()
}

func InsertReviewTransaction(db *sql.DB, review model.PostReviewTransaction, transactionId int) (err error) {
	sql := "UPDATE transactions SET rating = $1, review = $2, updated_at = $3 WHERE id = $4;"
	res, err := db.Exec(sql, review.Rating, review.Review, time.Now().Local(), transactionId)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Failed to delete data because transaction data is not found")
	}
	return nil
}

func ProceedTransaction(db *sql.DB, transactionId int) (err error) {
	sql := "UPDATE transactions SET status = $1, updated_at = $2 WHERE id = $3;"
	res, err := db.Exec(sql, "In Progress", time.Now().Local(), transactionId)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Failed to delete data because transaction data is not found")
	}
	return nil
}

func CancelTransaction(db *sql.DB, transactionId int) (err error) {
	sql := "UPDATE transactions SET status = $1, updated_at = $2 WHERE id = $3;"
	res, err := db.Exec(sql, "Cancelled", time.Now().Local(), transactionId)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Failed to delete data because transaction data is not found")
	}
	return nil
}

func FinishTransaction(db *sql.DB, transactionId int) (err error) {
	sql := "UPDATE transactions SET status = $1, updated_at = $2 WHERE id = $3;"
	res, err := db.Exec(sql, "Finished", time.Now().Local(), transactionId)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Failed to delete data because transaction data is not found")
	}
	return nil
}

func GetTransactionStatus(db *sql.DB, id int) (string, error) {
	var status string
	sqlStatement := `SELECT status from transactions where id = $1`
	// Query for a value based on a single row.
	if err := db.QueryRow(sqlStatement, id).Scan(&status); err != nil {
		return "", errors.New("Transaction with that ID is not found in database")
	}
	return status, nil
}

func GetCustomerIdInTransaction(db *sql.DB, id int) (int, error) {
	var customerId int
	sqlStatement := `SELECT customer_id from transactions where id = $1`
	// Query for a value based on a single row.
	if err := db.QueryRow(sqlStatement, id).Scan(&customerId); err != nil {
		return 0, errors.New("Transaction with that ID is not found in database")
	}
	return customerId, nil
}
