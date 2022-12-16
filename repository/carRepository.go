package repository

import (
	"car-rental/model"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

func GetAllCar(db *sql.DB) (results []model.Car, err error) {
	sql := "SELECT cars.id, cars.car_number, cars.car_model, cars.car_year, cars.rent_price, cars.brand_id, brands.name AS brand_name, cars.created_at, cars.updated_at FROM cars INNER JOIN brands on cars.brand_id = brands.id"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var car = model.Car{}
		err := rows.Scan(&car.ID, &car.CarNumber, &car.CarModel, &car.CarYear, &car.RentPrice, &car.BrandId, &car.CarBrand, &car.CreatedAt, &car.UpdatedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, car)
	}
	return
}

func InsertCar(db *sql.DB, car model.PostPutCar) (err error) {
	sql := "INSERT INTO cars (car_number, car_model, car_year, rent_price, brand_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	errs := db.QueryRow(sql, car.CarNumber, car.CarModel, car.CarYear, car.RentPrice, car.BrandId, time.Now().Local(), time.Now().Local())
	return errs.Err()
}

func UpdateCar(db *sql.DB, car model.PostPutCar, carId int) (err error) {
	sql := "UPDATE car SET car_number = $1, car_model = $2, car_year = $3, rent_price = $4, brand_id = $5, updated_at = $6 WHERE id = $7;"
	res, err := db.Exec(sql, car.CarNumber, car.CarModel, car.CarYear, car.RentPrice, car.BrandId, time.Now().Local(), carId)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Failed to delete data because car data is not found")
	}
	return nil
}

func DeleteCar(db *sql.DB, car model.Car) (err error) {
	sql := "DELETE from cars WHERE id = $1;"
	res, err := db.Exec(sql, car.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Failed to delete data because car data is not found")
	}
	return nil
}

func GetCarPrice(db *sql.DB, id int) (int, error) {
	var price int
	sqlStatement := `SELECT rent_price from cars where id = $1`
	// Query for a value based on a single row.
	if err := db.QueryRow(sqlStatement, id).Scan(&price); err != nil {
		return 0, errors.New("Car with that ID is not found in database")
	}
	return price, nil
}

func IsCarCanBooked(db *sql.DB, CarId int, dateStart string, dateFinish string) error {
	var id int
	sqlStatement := `SELECT id from transactions WHERE car_id = $1 AND (status = $2 OR status = $3) AND ( $4 >= date_start AND $5 <= date_finish )`
	if err := db.QueryRow(sqlStatement, CarId, "Created", "In Progress", dateStart, dateStart).Scan(&id); err != nil {
		fmt.Println(id)
		return nil
	}
	fmt.Println(id)
	return errors.New("Transaction cannot happen because the car is already booked")
}
