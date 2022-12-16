package route

import (
	"car-rental/controllers"
	"car-rental/middleware"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	router := gin.Default()

	//Auth
	router.POST("/admins/login", controllers.AuthLoginAdmin)
	router.POST("/customers/login", controllers.AuthLoginCustomer)

	//Admin
	router.GET("/admins", middleware.JwtAuthMiddleware(), controllers.GetAllAdmin)
	router.POST("/admins", controllers.InsertAdmin)
	router.PUT("/admins/:id", middleware.JwtAuthMiddleware(), controllers.UpdateAdmin)
	router.DELETE("/admins/:id", middleware.JwtAuthMiddleware(), controllers.DeleteAdmin)

	//Customer
	router.POST("/customers", controllers.InsertCustomer)
	router.GET("/customers", middleware.JwtAuthMiddleware(), controllers.GetAllCustomer)
	router.PUT("/customers/:id", middleware.JwtAuthMiddleware(), controllers.UpdateCustomer)
	router.DELETE("/customers/:id", middleware.JwtAuthMiddleware(), controllers.DeleteCustomer)

	//Brand
	router.POST("/brands", middleware.JwtAuthMiddleware(), controllers.InsertBrand)
	router.GET("/brands", controllers.GetAllBrand)
	router.PUT("/brands/:id", middleware.JwtAuthMiddleware(), controllers.UpdateBrand)
	router.DELETE("/brands/:id", middleware.JwtAuthMiddleware(), controllers.DeleteBrand)

	//Car
	router.POST("/cars", middleware.JwtAuthMiddleware(), controllers.InsertCar)
	router.GET("/cars", controllers.GetAllCar)
	router.PUT("/cars/:id", middleware.JwtAuthMiddleware(), controllers.UpdateCar)
	router.DELETE("/cars/:id", middleware.JwtAuthMiddleware(), controllers.DeleteCar)
	router.GET("/cars/:id/car_images", controllers.GetCarImageByCarId)

	//Car Image
	router.POST("/car_images", middleware.JwtAuthMiddleware(), controllers.InsertCarImage)
	router.PUT("/car_images/:id", middleware.JwtAuthMiddleware(), controllers.UpdateCarImage)
	router.DELETE("/car_images/:id", middleware.JwtAuthMiddleware(), controllers.DeleteCarImage)

	//Transaction
	router.POST("/transactions", middleware.JwtAuthMiddleware(), controllers.InsertTransaction)
	router.GET("/transactions", middleware.JwtAuthMiddleware(), controllers.GetAllTransaction)
	router.GET("/cars/:id/transactions", middleware.JwtAuthMiddleware(), controllers.GetTransactionByCarId)
	router.GET("/customers/:id/transactions", middleware.JwtAuthMiddleware(), controllers.GetTransactionByCustomerId)
	router.PUT("/transactions/:id/proceed", middleware.JwtAuthMiddleware(), controllers.ProceedTransaction)
	router.PUT("/transactions/:id/finished", middleware.JwtAuthMiddleware(), controllers.FinishTransaction)
	router.PUT("/transactions/:id/cancelled", middleware.JwtAuthMiddleware(), controllers.CancelTransaction)
	router.PUT("/transactions/:id/review", middleware.JwtAuthMiddleware(), controllers.InsertReviewTransaction)

	return router
}
