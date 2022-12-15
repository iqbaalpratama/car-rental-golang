package repository

import (
	"car-rental/model"
	"database/sql"
	"time"
)

func GetAllBrand(db *sql.DB) (err error, results []model.Brand) {
	sql := "SELECT id, name, created_at, updated_at FROM brands"
	rows, err := db.Query(sql)
	if err != nil {
		return err, nil
	}
	defer rows.Close()
	for rows.Next() {
		var brand = model.Brand{}
		err := rows.Scan(&brand.ID, &brand.Name, &brand.CreatedAt, &brand.UpdatedAt)
		if err != nil {
			return err, nil
		}
		results = append(results, brand)
	}
	return
}

func InsertBrand(db *sql.DB, brand model.PostPutBrand) (err error) {
	sql := "INSERT INTO brands (name, created_at, updated_at) VALUES ($1, $2, $3)"
	errs := db.QueryRow(sql, brand.Name, time.Now().Local(), time.Now().Local())
	return errs.Err()
}

// func UpdatePerson(db *sql.DB, person model.Admin) (err error) {
// 	sql := "UPDATE person SET first_name = $1, last_name = $2 WHERE id = $3"
// 	errs := db.QueryRow(sql, person.FirstName, person.LastName, person.ID)
// 	return errs.Err()
// }

func DeleteBrand(db *sql.DB, brand model.Brand) (err error) {
	sql := "DELETE from brands WHERE id = $1"
	errs := db.QueryRow(sql, brand.ID)
	return errs.Err()
}
