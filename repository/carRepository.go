package repository

import (
	"car-rental/model"
	"database/sql"
	"time"
)

func GetAllCar(db *sql.DB) (err error, results []model.Car) {
	sql := "SELECT id, car_number, car_model, car_year, rent_price, brand_id, brands.name AS brand_name, created_at, updated_at FROM cars INNER JOIN brands on cars.brand_id = brands.id"
	rows, err := db.Query(sql)
	if err != nil {
		return err, nil
	}
	defer rows.Close()
	for rows.Next() {
		var car = model.Car{}
		err := rows.Scan(&car.ID, &car.CarNumber, &car.CarModel, &car.CarYear, &car.RentPrice, &car.BrandId, &car.CarBrand, &car.CreatedAt, &car.UpdatedAt)
		if err != nil {
			return err, nil
		}
		results = append(results, car)
	}
	return
}

func InsertCar(db *sql.DB, car model.PostCar) (err error) {
	sql := "INSERT INTO cars (car_number, car_model, car_year, rent_price, brand_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	errs := db.QueryRow(sql, car.CarNumber, car.CarModel, car.CarYear, car.RentPrice, car.BrandId, time.Now().Local(), time.Now().Local())
	return errs.Err()
}

// func UpdatePerson(db *sql.DB, person model.Admin) (err error) {
// 	sql := "UPDATE person SET first_name = $1, last_name = $2 WHERE id = $3"
// 	errs := db.QueryRow(sql, person.FirstName, person.LastName, person.ID)
// 	return errs.Err()
// }

func DeleteCar(db *sql.DB, car model.Car) (err error) {
	sql := "DELETE from cars WHERE id = $1"
	errs := db.QueryRow(sql, car.ID)
	return errs.Err()
}
