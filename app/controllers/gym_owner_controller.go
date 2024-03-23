package controllers

import (
	"kalorize-api/app/services"
	"strings"

	vl "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type GymOwnerController struct {
	gymOwnerService services.GymOwnerService
	validate        vl.Validate
}

func NewGymOwnerController(db *gorm.DB) GymOwnerController {
	service := services.NewGymOwnerService(db)
	controller := GymOwnerController{
		gymOwnerService: service,
		validate:        *vl.New(),
	}
	return controller
}

func (controller *GymOwnerController) GenerateKodeGym(c echo.Context) error {

	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	type payload struct {
		IdGym uuid.UUID `json:"idGym" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	response := controller.gymOwnerService.GenerateKodeGym(payloadValidator.IdGym)
	return c.JSON(response.StatusCode, response)
}
