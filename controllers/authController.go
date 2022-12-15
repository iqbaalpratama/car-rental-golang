package controllers

import (
	"car-rental/database"
	"car-rental/model"
	"car-rental/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthLoginAdmin(c *gin.Context) {
	var input model.AuthLogin

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := repository.LoginCheckAdmin(database.DbConnection, input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func AuthLoginCustomer(c *gin.Context) {
	var input model.AuthLogin

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := repository.LoginCheckCustomer(database.DbConnection, input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
