package main

import (
	"log"
	"os"

	"github.com/D0T7/url_shortner/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	setupRouter(router)
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(router.Run(":" + port))
}

func setupRouter(router *gin.Engine) {
	router.POST("/api/v1/shorten", routes.ShortenUrl)
}
