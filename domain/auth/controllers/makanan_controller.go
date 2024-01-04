package controllers

import (
	"fmt"
	"kalorize-api/domain/auth/services"
	"strings"

	vl "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MakananController struct {
	makananService services.MakananService
	validate       vl.Validate
}

func NewMakananController(db *gorm.DB) MakananController {
	service := services.NewMakananService(db)
	controller := MakananController{
		makananService: service,
		validate:       *vl.New(),
	}
	return controller
}

func (controller *MakananController) GetAllMakanan(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	fmt.Print("masuk controller")
	response := controller.makananService.GetAllMakanan()
	return c.JSON(response.StatusCode, response)
}

func (controller *MakananController) GetMakananById(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	fmt.Print("Param: ", c.Param("makananId"))
	response := controller.makananService.GetMakananById(c.Param("makananId"))
	return c.JSON(response.StatusCode, response)
}
