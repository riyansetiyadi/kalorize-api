package main

import (
	"fmt"
	"kalorize-api/config"
	"kalorize-api/domain/auth/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := config.InitDB()

	apiv1 := e.Group("/api/v1")

	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Token{})

	authController := controllers.NewAuthController(db)

	apiv1.POST("/login", authController.Login)
	apiv1.POST("/register", authController.Register)
	apiv1.GET("/user", authController.GetUser)

	// Start server
	port := 8080
	address := fmt.Sprintf(":%d", port)
	e.Logger.Fatal(e.Start(address))
}
