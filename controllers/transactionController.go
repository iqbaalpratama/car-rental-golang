package controllers

import (
	"car-rental/database"
	"car-rental/model"
	"car-rental/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTransaction(c *gin.Context) {
	var (
		result gin.H
	)
	transactions, err := repository.GetAllTransaction(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": transactions,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertTransaction(c *gin.Context) {
	var transaction model.PostTransaction
	err := c.ShouldBindJSON(&transaction)
	if err != nil {
		panic(err)
	}
	err = repository.InsertTransaction(database.DbConnection, transaction)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Transaction",
	})
}

func InsertReviewTransaction(c *gin.Context) {
	var review model.PostReviewTransaction
	err := c.ShouldBindJSON(&review)
	if err != nil {
		panic(err)
	}
	err = repository.InsertReviewTransaction(database.DbConnection, review)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Transaction",
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

func DeleteTransaction(c *gin.Context) {
	var transaction model.Transaction
	id, err := strconv.Atoi(c.Param("id"))
	transaction.ID = id
	err = repository.DeleteTransaction(database.DbConnection, transaction)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Transaction",
	})
}
