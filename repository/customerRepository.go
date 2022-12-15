package repository

import (
	"car-rental/model"
	"database/sql"
	"time"
)

func GetAllCustomer(db *sql.DB) (err error, results []model.Customer) {
	sql := "SELECT id, first_name, last_name, email, address, phone_number, created_at, updated_at FROM customers"
	rows, err := db.Query(sql)
	if err != nil {
		return err, nil
	}
	defer rows.Close()
	for rows.Next() {
		var customer = model.Customer{}
		err := rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Address, &customer.PhoneNumber, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			return err, nil
		}
		results = append(results, customer)
	}
	return
}

func InsertCustomer(db *sql.DB, admin model.PostCustomer) (err error) {
	sql := "INSERT INTO customers (first_name, last_name, email, password, address, phone_number, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	errs := db.QueryRow(sql, admin.FirstName, admin.LastName, admin.Email, admin.Password, admin.Address, admin.PhoneNumber, time.Now().Local(), time.Now().Local())
	return errs.Err()
}

// func UpdatePerson(db *sql.DB, person model.Admin) (err error) {
// 	sql := "UPDATE person SET first_name = $1, last_name = $2 WHERE id = $3"
// 	errs := db.QueryRow(sql, person.FirstName, person.LastName, person.ID)
// 	return errs.Err()
// }

func DeleteCustomer(db *sql.DB, admin model.Customer) (err error) {
	sql := "DELETE from customers WHERE id = $1"
	errs := db.QueryRow(sql, admin.ID)
	return errs.Err()
}
