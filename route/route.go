package route

import (
	"car-rental/controllers"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	router := gin.Default()
	router.GET("/admins", controllers.GetAllAdmin)
	router.POST("/admins", controllers.InsertAdmin)
	// router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/admins/:id", controllers.DeleteAdmin)
	router.GET("/customers", controllers.GetAllCustomer)
	router.POST("/customers", controllers.InsertCustomer)
	// router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/customers/:id", controllers.DeleteCustomer)
	router.GET("/brands", controllers.GetAllBrand)
	router.POST("/brands", controllers.InsertBrand)
	// router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/brands/:id", controllers.DeleteBrand)
	return router
}
