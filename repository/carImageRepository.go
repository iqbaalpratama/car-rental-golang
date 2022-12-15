package repository

import (
	"car-rental/model"
	"database/sql"
	"errors"
	"time"
)

func GetCarImagesByCarId(db *sql.DB, carImage model.CarImage) (results []model.CarImage, err error) {
	sqlStatement := `SELECT * from car_images where car_id = $1`
	rows, err := db.Query(sqlStatement, carImage.CarId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var carImage = model.CarImage{}
		err = rows.Scan(&carImage.ID, &carImage.ImageURL, &carImage.CarId, &carImage.CreatedAt, &carImage.UpdatedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, carImage)
	}
	return
}

func InsertCarImage(db *sql.DB, carImage model.CarImage) (err error) {
	sql := "INSERT INTO car_images(url_image, car_id, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	errs := db.QueryRow(sql, carImage.ImageURL, carImage.CarId, time.Now().Local(), time.Now().Local())
	return errs.Err()
}

func UpdateCarImage(db *sql.DB, carImage model.CarImage) (err error) {
	sql := "UPDATE car_images SET image_url = $1, car_id = $2, updated_at = $3 WHERE id = $4;"
	res, err := db.Exec(sql, carImage.ImageURL, carImage.CarId, carImage.UpdatedAt, carImage.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Failed to delete data because car image data is not found")
	}
	return nil
}

func DeleteCarImage(db *sql.DB, carImage model.CarImage) (err error) {
	sql := "DELETE from car_images WHERE id = $1;"
	res, err := db.Exec(sql, carImage.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Failed to delete data because car image is not found")
	}
	return nil
}
