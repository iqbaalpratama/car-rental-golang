package repository

import (
	"car-rental/model"
	"database/sql"
	"time"
)

func GetCarImageByCarId(db *sql.DB, carImage model.CarImage) (err error, results []model.CarImage) {
	sql := "SELECT id, url_image, car_id, created_at, updated_at FROM car_images where car_id = $1"
	rows, err := db.Query(sql, carImage.CarId)
	if err != nil {
		return err, nil
	}
	defer rows.Close()
	for rows.Next() {
		var carImage = model.CarImage{}
		err := rows.Scan(&carImage.ID, &carImage.ImageURL, &carImage.CarId, &carImage.CreatedAt, &carImage.UpdatedAt)
		if err != nil {
			return err, nil
		}
		results = append(results, carImage)
	}
	return
}

func InsertCarImage(db *sql.DB, carImage model.PostCarImage) (err error) {
	sql := "INSERT INTO car_images(url_image, car_id, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	errs := db.QueryRow(sql, carImage.ImageURL, carImage.CarId, time.Now().Local(), time.Now().Local())
	return errs.Err()
}

// func UpdatePerson(db *sql.DB, person model.Admin) (err error) {
// 	sql := "UPDATE person SET first_name = $1, last_name = $2 WHERE id = $3"
// 	errs := db.QueryRow(sql, person.FirstName, person.LastName, person.ID)
// 	return errs.Err()
// }

func DeleteCarImage(db *sql.DB, carImage model.CarImage) (err error) {
	sql := "DELETE from car_images WHERE id = $1"
	errs := db.QueryRow(sql, carImage.ID)
	return errs.Err()
}
