package routes

import (
	"kalorize-api/domain/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteUser(apiv1 *echo.Group, db *gorm.DB) {
	userController := controllers.NewUserController(db)

	apiv1.PUT("/edit-user", userController.EditUser)
	apiv1.PUT("/edit-password", userController.EditPassword)
	apiv1.PUT("/edit-photo", userController.EditPhoto)
	apiv1.POST("/user/history", userController.CreateHistory)
	apiv1.GET("/user/history", userController.GetHistoryBaseDateTime)
}
