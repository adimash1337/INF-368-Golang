package main

import (
	"log"
	"os"

	"github.com/adimash1337/GoMid/middleware"
	"github.com/adimash1337/GoMid/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	log.Fatal(router.Run(":" + port))
}
