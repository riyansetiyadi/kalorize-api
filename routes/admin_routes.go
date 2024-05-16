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
	apiv1.POST("	", adminController.RegisterUser)
	apiv1.GET("/admin/get-all-user", adminController.GetAllUser)
	apiv1.GET("/admin/get-user/:id", adminController.GetUserById)
	apiv1.PUT("/admin/update-user/:id", adminController.UpdateUser)
	apiv1.DELETE("/admin/delete-user/:id", adminController.DeleteUser)
}
