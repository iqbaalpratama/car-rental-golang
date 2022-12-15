package controllers

import (
	"car-rental/database"
	"car-rental/helper"
	"car-rental/model"
	"car-rental/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCustomer(c *gin.Context) {
	var (
		result gin.H
	)
	customers, err := repository.GetAllCustomer(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": customers,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertCustomer(c *gin.Context) {
	var customer model.PostCustomer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		panic(err)
	}
	password, errPassword := helper.HashPassword(customer.Password)
	fmt.Println(password)
	if errPassword != nil {
		panic(errPassword)
	}
	customer.Password = password
	err = repository.InsertCustomer(database.DbConnection, customer)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Customer",
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

func DeleteCustomer(c *gin.Context) {
	var customer model.Customer
	id, err := strconv.Atoi(c.Param("id"))
	customer.ID = id
	err = repository.DeleteCustomer(database.DbConnection, customer)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Admin",
	})
}
