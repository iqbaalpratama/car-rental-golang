package repository

import (
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

// func GetAdminById(db *sql.DB) (err error, admin model.Admin) {
// 	sql := "SELECT id, first_name, last_name, email, phone_number, created_at, updated_at FROM admins"
// 	rows, err := db.Query(sql)
// 	if err != nil {
// 		return err, nil
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var admin = model.Admin{}
// 		err := rows.Scan(&admin.ID, &admin.FirstName, &admin.LastName, &admin.Email, &admin.PhoneNumber, &admin.CreatedAt, &admin.UpdatedAt)
// 		if err != nil {
// 			return err, nil
// 		}
// 		results = append(results, admin)
// 	}
// 	return
// }

func InsertAdmin(db *sql.DB, admin model.PostPutAdmin) (err error) {
	sql := "INSERT INTO admins (first_name, last_name, email, password, phone_number, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	errs := db.QueryRow(sql, admin.FirstName, admin.LastName, admin.Email, admin.Password, admin.PhoneNumber, time.Now().Local(), time.Now().Local())
	return errs.Err()
}

// func UpdatePerson(db *sql.DB, person model.Admin) (err error) {
// 	sql := "UPDATE person SET first_name = $1, last_name = $2 WHERE id = $3"
// 	errs := db.QueryRow(sql, person.FirstName, person.LastName, person.ID)
// 	return errs.Err()
// }

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
