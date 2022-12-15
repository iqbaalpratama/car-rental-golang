package controllers

import (
	"car-rental/database"
	"car-rental/helper"
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
	var car model.PostPutCar
	var errorValidation []string
	err := c.ShouldBindJSON(&car)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !helper.IsCarYearValid(car.CarYear) {
		errorValidation = append(errorValidation, "Car year is not valid, it should between 2006 and 2022")
	}
	if len(errorValidation) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Input data is not valid",
			"error_message": errorValidation,
		})
		return
	}
	err = repository.InsertCar(database.DbConnection, car)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Car",
	})
}

func UpdateCar(c *gin.Context) {
	var car model.PostPutCar
	var errorValidation []string
	carId, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&car)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !helper.IsCarYearValid(car.CarYear) {
		errorValidation = append(errorValidation, "Car year is not valid, it should between 2006 and 2022")
	}
	if len(errorValidation) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Input data is not valid",
			"error_message": errorValidation,
		})
		return
	}
	err = repository.UpdateCar(database.DbConnection, car, carId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Car",
	})
}

func DeleteCar(c *gin.Context) {
	var car model.Car
	id, err := strconv.Atoi(c.Param("id"))
	car.ID = id
	err = repository.DeleteCar(database.DbConnection, car)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Car",
	})
}
