package controllers

import (
	"car-rental/database"
	"car-rental/model"
	"car-rental/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllBrand(c *gin.Context) {
	var (
		result gin.H
	)
	brands, err := repository.GetAllBrand(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": brands,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertBrand(c *gin.Context) {
	var brand model.Brand
	err := c.ShouldBindJSON(&brand)
	if err != nil {
		panic(err)
	}
	brand.CreatedAt = time.Now().Local()
	brand.UpdatedAt = time.Now().Local()
	err = repository.InsertBrand(database.DbConnection, brand)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Brand",
	})
}

func UpdateBrand(c *gin.Context) {
	var brand model.Brand
	brandId, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&brand)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	brand.ID = brandId
	brand.UpdatedAt = time.Now().Local()
	err = repository.UpdateBrand(database.DbConnection, brand)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Brand",
	})
}

func DeleteBrand(c *gin.Context) {
	var brand model.Brand
	id, err := strconv.Atoi(c.Param("id"))
	brand.ID = id
	err = repository.DeleteBrand(database.DbConnection, brand)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Brand",
	})
}
