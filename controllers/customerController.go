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
	var errorValidation []string
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !helper.IsLengthPasswordValid(customer.Password) {
		errorValidation = append(errorValidation, "Password length is not valid, it should contain minimum 6 characters")
	}
	if !helper.IsEmailValid(customer.Email) {
		errorValidation = append(errorValidation, "Email input is not valid")
	}
	if !helper.IsLengthPhoneNumberValid(customer.PhoneNumber) {
		errorValidation = append(errorValidation, "Phone number length is not valid, it should be between 11 and 14 characters length")
	}
	if len(errorValidation) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Input data is not valid",
			"error_message": errorValidation,
		})
		return
	}

	password, err := helper.HashPassword(customer.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Admin",
	})
}
