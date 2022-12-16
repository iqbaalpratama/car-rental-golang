package repository

import (
	"car-rental/helper"
	"car-rental/model"
	"database/sql"
	"errors"
	"time"
)

func GetAllAdmin(db *sql.DB) (results []model.Admin, err error) {
	sql := "SELECT id, first_name, last_name, email, phone_number, created_at, updated_at FROM admins"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var admin = model.Admin{}
		err := rows.Scan(&admin.ID, &admin.FirstName, &admin.LastName, &admin.Email, &admin.PhoneNumber, &admin.CreatedAt, &admin.UpdatedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, admin)
	}
	return
}

func InsertAdmin(db *sql.DB, admin model.PostAdmin) (err error) {
	sql := "INSERT INTO admins (first_name, last_name, email, password, phone_number, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	errs := db.QueryRow(sql, admin.FirstName, admin.LastName, admin.Email, admin.Password, admin.PhoneNumber, time.Now().Local(), time.Now().Local())
	return errs.Err()
}

func UpdateAdmin(db *sql.DB, admin model.PutAdmin, adminId int) (err error) {
	sql := "UPDATE admins SET first_name = $1, last_name = $2, phone_number = $3, updated_at = $4 WHERE id = $5;"
	res, err := db.Exec(sql, admin.FirstName, admin.LastName, admin.PhoneNumber, time.Now().Local(), adminId)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Failed to update data because admin data is not found")
	}
	return nil
}

func DeleteAdmin(db *sql.DB, admin model.Admin) (err error) {
	sql := "DELETE from admins WHERE id = $1;"
	res, err := db.Exec(sql, admin.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Failed to delete data because admin data is not found")
	}
	return nil
}

func LoginCheckAdmin(db *sql.DB, input model.AuthLogin) (string, error) {
	var id int
	var password string
	sql := "SELECT id, password FROM admins WHERE email = $1"
	if err := db.QueryRow(sql, input.Email).Scan(&id, &password); err != nil {
		return "", errors.New("Email not found")
	}
	if !helper.CheckPasswordHash(input.Password, password) {
		return "", errors.New("Password not match with email")
	}
	data := model.Token{ID: id, Role: "Admin"}
	token, err := helper.GenerateToken(data)
	if err != nil {
		return "", err
	}
	return token, nil
}
