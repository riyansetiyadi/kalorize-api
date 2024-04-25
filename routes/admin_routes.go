package routes

import (
	"kalorize-api/app/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RoutesAdmin(apiv1 *echo.Group, db *gorm.DB) {
	adminController := controllers.NewAdminController(db)

	apiv1.POST("/admin/create-makanan", adminController.RegisterMakanan)
	apiv1.POST("/admin/create-gym", adminController.RegisterGym)
	apiv1.POST("/admin/create-franchise", adminController.RegisterFranchise)
	apiv1.POST("/admin/create-gymcode", adminController.GenerateGymToken)
}
