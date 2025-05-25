package main

import (
	"go-mvc-demo/config"
	routes "go-mvc-demo/router"
	"log"
	"os"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go-mvc-demo/docs" // **Bắt buộc phải import package docs**
)

// @title Demo Swagger API
// @version 1.0
// @description  Swagger social media
// @host localhost:8080
// @BasePath

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.ConnectDatabase()
	r := routes.SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + port)
}
