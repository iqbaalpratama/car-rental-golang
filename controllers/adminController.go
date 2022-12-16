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
	var errorValidation []string
	var admin model.PostAdmin
	err := c.ShouldBindJSON(&admin)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !helper.IsLengthPasswordValid(admin.Password) {
		errorValidation = append(errorValidation, "Password length is not valid, it should contain minimum 6 characters")
	}
	if !helper.IsEmailValid(admin.Email) {
		errorValidation = append(errorValidation, "Email input is not valid")
	}
	if !helper.IsLengthPhoneNumberValid(admin.PhoneNumber) {
		errorValidation = append(errorValidation, "Phone number length is not valid, it should be between 11 and 14 characters length")
	}
	if len(errorValidation) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Input data is not valid",
			"error_message": errorValidation,
		})
		return
	}

	password, err := helper.HashPassword(admin.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	admin.Password = password

	err = repository.InsertAdmin(database.DbConnection, admin)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Admin",
	})
}

func UpdateAdmin(c *gin.Context) {
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
	var admin model.PutAdmin
	var errorValidation []string
	adminId, _ := strconv.Atoi(c.Param("id"))
	if dataUser.ID != adminId {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "You are unauthorized to access this resource, this is for another user",
		})
		return
	}
	err = c.ShouldBindJSON(&admin)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !helper.IsLengthPhoneNumberValid(admin.PhoneNumber) {
		errorValidation = append(errorValidation, "Phone number length is not valid, it should be between 11 and 14 characters length")
	}
	if len(errorValidation) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Input data is not valid",
			"error_message": errorValidation,
		})
		return
	}
	err = repository.UpdateAdmin(database.DbConnection, admin, adminId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Admin",
	})
}

func DeleteAdmin(c *gin.Context) {
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
	var admin model.Admin
	id, err := strconv.Atoi(c.Param("id"))
	admin.ID = id
	err = repository.DeleteAdmin(database.DbConnection, admin)
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
