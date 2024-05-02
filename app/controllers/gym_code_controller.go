package controllers

import (
	"kalorize-api/app/services"
	"strings"

	vl "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type GymController struct {
	gymService *services.GymService
	validate   vl.Validate
}

func NewGymController(db *gorm.DB) GymController {
	service := services.NewGymService(db)
	controller := GymController{
		gymService: service,
		validate:   *vl.New(),
	}
	return controller
}

func (controller *GymController) CheckGymCode(c echo.Context) error {
	type payload struct {
		GymCode string `json:"gym_code" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	response := controller.gymService.CheckGymCode(payloadValidator.GymCode)
	return c.JSON(response.StatusCode, response)
}

func (controller *GymController) IsUsed(c echo.Context) error {
	type payload struct {
		GymCode string `json:"gym_code" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	response := controller.gymService.IsUsed(payloadValidator.GymCode)
	return c.JSON(response.StatusCode, response)
}

// Get All Gym
func (controller *GymController) GetAllGym(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	response := controller.gymService.GetAllGym()
	return c.JSON(response.StatusCode, response)
}
