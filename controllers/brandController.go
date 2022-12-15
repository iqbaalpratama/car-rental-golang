package controllers

import (
	"car-rental/database"
	"car-rental/model"
	"car-rental/repository"
	"net/http"
	"strconv"

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
	var brand model.PostPutBrand
	err := c.ShouldBindJSON(&brand)
	if err != nil {
		panic(err)
	}
	err = repository.InsertBrand(database.DbConnection, brand)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Brand",
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

func DeleteBrand(c *gin.Context) {
	var brand model.Brand
	id, err := strconv.Atoi(c.Param("id"))
	brand.ID = id
	err = repository.DeleteBrand(database.DbConnection, brand)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Brand",
	})
}
