package routes

import (
	"kalorize-api/domain/auth/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteAuth(apiv1 *echo.Group, db *gorm.DB) {
	authController := controllers.NewAuthController(db)

	apiv1.POST("/login", authController.Login)
	apiv1.POST("/register", authController.Register)
	apiv1.GET("/user", authController.GetUser)
}
