package controllers

import (
	"car-rental/database"
	"car-rental/helper"
	"car-rental/model"
	"car-rental/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetCarImageByCarId(c *gin.Context) {
	var (
		result gin.H
	)
	var carImage model.CarImage
	id, err := strconv.Atoi(c.Param("id"))
	carImage.CarId = id
	carImages, err := repository.GetCarImagesByCarId(database.DbConnection, carImage)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": carImages,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertCarImage(c *gin.Context) {
	dataUser, err := helper.ExtractTokenData(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if dataUser == (model.Token{}) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Cannot extract data token",
		})
		return
	}
	if dataUser.Role == "Customer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "You are unauthorized to access this resource, this resource for admin user",
		})
		return
	}
	var carImage model.CarImage
	var errorValidation []string
	err = c.ShouldBindJSON(&carImage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !helper.IsUrlValid(carImage.ImageURL) {
		errorValidation = append(errorValidation, "Image URL not valid")
	}
	if len(errorValidation) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Input data is not valid",
			"error_message": errorValidation,
		})
		return
	}
	carImage.CreatedAt = time.Now().Local()
	carImage.UpdatedAt = time.Now().Local()
	err = repository.InsertCarImage(database.DbConnection, carImage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Car Image",
	})
}

func UpdateCarImage(c *gin.Context) {
	dataUser, err := helper.ExtractTokenData(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if dataUser == (model.Token{}) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Cannot extract data token",
		})
		return
	}
	if dataUser.Role == "Customer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "You are unauthorized to access this resource, this resource for admin user",
		})
		return
	}
	var carImage model.CarImage
	var errorValidation []string
	carImageId, _ := strconv.Atoi(c.Param("id"))
	err = c.ShouldBindJSON(&carImage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !helper.IsUrlValid(carImage.ImageURL) {
		errorValidation = append(errorValidation, "Image URL not valid")
	}
	if len(errorValidation) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Input data is not valid",
			"error_message": errorValidation,
		})
		return
	}
	carImage.ID = carImageId
	carImage.UpdatedAt = time.Now().Local()
	err = repository.UpdateCarImage(database.DbConnection, carImage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Car Image",
	})
}

func DeleteCarImage(c *gin.Context) {
	dataUser, err := helper.ExtractTokenData(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if dataUser == (model.Token{}) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Cannot extract data token",
		})
		return
	}
	if dataUser.Role == "Customer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "You are unauthorized to access this resource, this resource for admin user",
		})
		return
	}
	var carImage model.CarImage
	id, err := strconv.Atoi(c.Param("id"))
	carImage.ID = id
	err = repository.DeleteCarImage(database.DbConnection, carImage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Car Image",
	})
}
