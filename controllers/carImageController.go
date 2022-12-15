package controllers

import (
	"car-rental/database"
	"car-rental/model"
	"car-rental/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// func GetCarImageByCarId(c *gin.Context) {
// 	var (
// 		result gin.H
// 	)
// 	var carImage model.CarImage
// 	id, err := strconv.Atoi(c.Param("id"))
// 	carImage.CarId = id
// 	carImages, err := repository.GetCarImageByCarId(database.DbConnection, carImage)
// 	if err != nil {
// 		result = gin.H{
// 			"result": err,
// 		}
// 	} else {
// 		result = gin.H{
// 			"result": carImages,
// 		}
// 	}
// 	c.JSON(http.StatusOK, result)
// }

func InsertCarImage(c *gin.Context) {
	var carImage model.PostCarImage
	err := c.ShouldBindJSON(&carImage)
	if err != nil {
		panic(err)
	}
	err = repository.InsertCarImage(database.DbConnection, carImage)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Car Image",
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

func DeleteCarImage(c *gin.Context) {
	var carImage model.CarImage
	id, err := strconv.Atoi(c.Param("id"))
	carImage.ID = id
	err = repository.DeleteCarImage(database.DbConnection, carImage)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Car Image",
	})
}
