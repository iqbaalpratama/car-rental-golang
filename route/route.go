package route

import (
	"car-rental/controllers"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	router := gin.Default()
	//Admin
	router.GET("/admins", controllers.GetAllAdmin)
	router.POST("/admins", controllers.InsertAdmin)
	// router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/admins/:id", controllers.DeleteAdmin)

	//Customer
	router.GET("/customers", controllers.GetAllCustomer)
	router.POST("/customers", controllers.InsertCustomer)
	// router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/customers/:id", controllers.DeleteCustomer)

	//Brand
	router.GET("/brands", controllers.GetAllBrand)
	router.POST("/brands", controllers.InsertBrand)
	router.PUT("/brands/:id", controllers.UpdateBrand)
	router.DELETE("/brands/:id", controllers.DeleteBrand)

	//Car
	router.GET("/cars", controllers.GetAllCar)
	router.POST("/cars", controllers.InsertCar)
	router.PUT("/cars/:id", controllers.UpdateCar)
	router.DELETE("/cars/:id", controllers.DeleteCar)

	//Car Image
	router.POST("/car_images", controllers.InsertCarImage)
	router.PUT("/car_images/:id", controllers.UpdateCarImage)
	router.DELETE("/car_images/:id", controllers.DeleteCarImage)
	return router
}
