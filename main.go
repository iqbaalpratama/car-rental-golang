package main

import (
	"car-rental/database"
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Failed to load file environment")
	} else {
		fmt.Println("Success to load file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DATABASE"))
	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}
	database.DbMigrate(DB)
	defer DB.Close()
	// router := gin.Default()
	// router.GET("/persons", controllers.GetAllPerson)
	// router.POST("/persons", controllers.InsertPerson)
	// router.PUT("/persons/:id", controllers.UpdatePerson)
	// router.DELETE("/persons/:id", controllers.DeletePerson)
	// router.Run()
}
