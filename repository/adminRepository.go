package repository

import (
	"car-rental/model"
	"database/sql"
	"time"
)

func GetAllAdmin(db *sql.DB) (err error, results []model.Admin) {
	sql := "SELECT id, first_name, last_name, email, phone_number, created_at, updated_at FROM admins"
	rows, err := db.Query(sql)
	if err != nil {
		return err, nil
	}
	defer rows.Close()
	for rows.Next() {
		var admin = model.Admin{}
		err := rows.Scan(&admin.ID, &admin.FirstName, &admin.LastName, &admin.Email, &admin.PhoneNumber, &admin.CreatedAt, &admin.UpdatedAt)
		if err != nil {
			return err, nil
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

// func UpdatePerson(db *sql.DB, person model.Admin) (err error) {
// 	sql := "UPDATE person SET first_name = $1, last_name = $2 WHERE id = $3"
// 	errs := db.QueryRow(sql, person.FirstName, person.LastName, person.ID)
// 	return errs.Err()
// }

func DeleteAdmin(db *sql.DB, admin model.Admin) (err error) {
	sql := "DELETE from admins WHERE id = $1"
	errs := db.QueryRow(sql, admin.ID)
	return errs.Err()
}
