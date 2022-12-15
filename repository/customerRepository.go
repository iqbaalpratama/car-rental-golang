package repository

import (
	"car-rental/helper"
	"car-rental/model"
	"database/sql"
	"errors"
	"time"
)

func GetAllCustomer(db *sql.DB) (results []model.Customer, err error) {
	sql := "SELECT id, first_name, last_name, email, address, phone_number, created_at, updated_at FROM customers"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var customer = model.Customer{}
		err := rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Address, &customer.PhoneNumber, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			return nil, err
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

func DeleteCustomer(db *sql.DB, customer model.Customer) (err error) {
	sql := "DELETE from customers WHERE id = $1"
	res, err := db.Exec(sql, customer.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Failed to delete data because customer data is not found")
	}
	return nil
}

func LoginCheckCustomer(db *sql.DB, input model.AuthLogin) (string, error) {
	var id int
	var password string
	sql := "SELECT id, password FROM customers WHERE email = $1"
	if err := db.QueryRow(sql, input.Email).Scan(&id, &password); err != nil {
		return "", errors.New("Email not found")
	}
	if !helper.CheckPasswordHash(input.Password, password) {
		return "", errors.New("Password not match with email")
	}
	data := model.Token{ID: id, Role: "Customer"}
	token, err := helper.GenerateToken(data)
	if err != nil {
		return "", err
	}
	return token, nil
}
