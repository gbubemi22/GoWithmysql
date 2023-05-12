package main

import (
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/gbubemi/mysql2/controllers"
	routes "github.com/gbubemi/mysql2/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	
	"log"
	"os"
)



func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	userRepo := controllers.New()

	routes.UserRoutes(router, userRepo)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	

	

	// router.Run(":" + port)
	router.Run(":8001")

}