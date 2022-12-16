package repository

import (
	"car-rental/model"
	"database/sql"
	"errors"
)

func GetAllBrand(db *sql.DB) (results []model.Brand, err error) {
	sql := "SELECT id, name, created_at, updated_at FROM brands"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var brand = model.Brand{}
		err := rows.Scan(&brand.ID, &brand.Name, &brand.CreatedAt, &brand.UpdatedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, brand)
	}
	return
}

func InsertBrand(db *sql.DB, brand model.Brand) (err error) {
	sql := "INSERT INTO brands (name, created_at, updated_at) VALUES ($1, $2, $3)"
	errs := db.QueryRow(sql, brand.Name, brand.CreatedAt, brand.UpdatedAt)
	return errs.Err()
}

func UpdateBrand(db *sql.DB, brand model.Brand) (err error) {
	sql := "UPDATE brands SET name = $1, updated_at = $2 WHERE id = $3;"
	res, err := db.Exec(sql, brand.Name, brand.UpdatedAt, brand.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Failed to update data because brand data is not found")
	}
	return nil
}

func DeleteBrand(db *sql.DB, brand model.Brand) (err error) {
	sql := "DELETE from brands WHERE id = $1;"
	res, err := db.Exec(sql, brand.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Failed to delete data because brand data is not found")
	}
	return nil
}
