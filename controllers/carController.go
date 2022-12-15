package controllers

import (
	"car-rental/database"
	"car-rental/model"
	"car-rental/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCar(c *gin.Context) {
	var (
		result gin.H
	)
	cars, err := repository.GetAllCar(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": cars,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertCar(c *gin.Context) {
	var car model.PostCar
	err := c.ShouldBindJSON(&car)
	if err != nil {
		panic(err)
	}
	err = repository.InsertCar(database.DbConnection, car)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Car",
	})
}

// func UpdatePerson(c *gin.Context) {
// 	var person model.Person
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	err := c.ShouldBindJSON(&person)
// 	if err != nil {
// 		panic(err)
// 	}
// 	person.ID = int64(id)
// 	err = repository.UpdatePerson(database.DbConnection, person)
// 	if err != nil {
// 		panic(err)
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"result": "Success Update Person",
// 	})
// }

func DeletCar(c *gin.Context) {
	var car model.Car
	id, err := strconv.Atoi(c.Param("id"))
	car.ID = id
	err = repository.DeleteCar(database.DbConnection, car)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Car",
	})
}
