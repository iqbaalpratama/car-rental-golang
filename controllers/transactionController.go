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

func GetAllTransaction(c *gin.Context) {
	dataUser, errToken := helper.ExtractTokenData(c)
	if errToken != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": errToken.Error(),
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

func GetTransactionByCarId(c *gin.Context) {
	dataUser, errToken := helper.ExtractTokenData(c)
	if errToken != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": errToken.Error(),
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
	carId, _ := strconv.Atoi(c.Param("id"))
	var (
		result gin.H
	)
	transactions, err := repository.GetTransactionByCarId(database.DbConnection, carId)
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

func GetTransactionByCustomerId(c *gin.Context) {
	dataUser, errToken := helper.ExtractTokenData(c)
	if errToken != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": errToken.Error(),
		})
		return
	}
	if dataUser == (model.Token{}) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Cannot extract data token",
		})
		return
	}
	customerId, _ := strconv.Atoi(c.Param("id"))
	if dataUser.Role == "Customer" && dataUser.ID != customerId {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "You are unauthorized to access this resource, this resource for another user",
		})
		return
	}
	var (
		result gin.H
	)
	transactions, err := repository.GetTransactionByCustomerId(database.DbConnection, customerId)
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
	var errorValidation []string
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
	var transaction model.PostTransaction
	err = c.ShouldBindJSON(&transaction)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if dataUser.Role == "Customer" && dataUser.ID != transaction.CustomerId {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "You are unauthorized to access this resource",
		})
		return
	}
	//format date yyyy/mm/dd
	isValid, differenceDate := helper.IsDateValid(transaction.DateStart, transaction.DateFinish)
	if !isValid {
		errorValidation = append(errorValidation, "Date start and date finish not valid, it should have difference of positive")
	}
	if len(errorValidation) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Input data is not valid",
			"error_message": errorValidation,
		})
		return
	}
	rentPrice, err := repository.GetCarPrice(database.DbConnection, transaction.CarId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	transaction.TotalPrice = rentPrice * differenceDate
	err = repository.InsertTransaction(database.DbConnection, transaction)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Transaction",
	})
}

func InsertReviewTransaction(c *gin.Context) {
	var review model.PostReviewTransaction
	var status string
	var customerId int
	var errorValidation []string
	dataUser, errToken := helper.ExtractTokenData(c)
	if errToken != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": errToken.Error(),
		})
		return
	}
	if dataUser == (model.Token{}) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Cannot extract data token",
		})
		return
	}
	if dataUser.Role == "Admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "You are unauthorized to access this resource, this resource for customer user",
		})
		return
	}
	err := c.ShouldBindJSON(&review)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !helper.IsRatingValid(review.Rating) {
		errorValidation = append(errorValidation, "Rating is not valid, it should between 1 - 10")
	}
	if len(errorValidation) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Input data is not valid",
			"error_message": errorValidation,
		})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	status, err = repository.GetTransactionStatus(database.DbConnection, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if status != "Finished" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "This transactions can't be finish because status transaction is not finished",
		})
		return
	}

	customerId, err = repository.GetCustomerIdInTransaction(database.DbConnection, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if dataUser.ID != customerId {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "You are unauthorized to access this resource, this resource for another user",
		})
		return
	}

	err = repository.InsertReviewTransaction(database.DbConnection, review, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Review Transaction",
	})
}

func ProceedTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	status, err := repository.GetTransactionStatus(database.DbConnection, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if status != "Created" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "This transactions can't be proceed because status transaction is not created",
		})
		return
	}
	err = repository.ProceedTransaction(database.DbConnection, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Proceed Transaction",
	})
}

func CancelTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	status, err := repository.GetTransactionStatus(database.DbConnection, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if status != "Created" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "This transactions can't be cancel because status transaction is not created",
		})
		return
	}
	err = repository.CancelTransaction(database.DbConnection, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Cancel Transaction",
	})
}

func FinishTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	status, err := repository.GetTransactionStatus(database.DbConnection, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if status != "In Progress" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "This transactions can't be finish because status transaction is not created",
		})
		return
	}
	err = repository.CancelTransaction(database.DbConnection, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success finish Transaction",
	})
}
