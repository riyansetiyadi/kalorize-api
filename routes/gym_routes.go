package routes

import (
	"kalorize-api/app/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteGym(apiv1 *echo.Group, db *gorm.DB) {
	gymController := controllers.NewGymController(db)

	apiv1.GET("/gym", gymController.GetAllGym)
	apiv1.POST("/gym/:id", gymController.CheckGymCode)
	apiv1.POST("/gym/used/:id", gymController.IsUsed)
}
