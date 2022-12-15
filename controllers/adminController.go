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

func GetAllAdmin(c *gin.Context) {
	var (
		result gin.H
	)
	admins, err := repository.GetAllAdmin(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": admins,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertAdmin(c *gin.Context) {
	var admin model.PostAdmin
	err := c.ShouldBindJSON(&admin)
	if err != nil {
		panic(err)
	}
	password, errPassword := helper.HashPassword(admin.Password)
	fmt.Println(password)
	if errPassword != nil {
		panic(errPassword)
	}
	admin.Password = password
	err = repository.InsertAdmin(database.DbConnection, admin)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Admin",
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

func DeleteAdmin(c *gin.Context) {
	var admin model.Admin
	id, err := strconv.Atoi(c.Param("id"))
	admin.ID = id
	err = repository.DeleteAdmin(database.DbConnection, admin)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Admin",
	})
}
