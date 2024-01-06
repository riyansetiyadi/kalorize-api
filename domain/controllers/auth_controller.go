package controllers

import (
	"fmt"
	"kalorize-api/domain/services"
	"kalorize-api/utils"
	"strings"

	vl "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthController struct {
	authService services.AuthService
	validate    vl.Validate
}

func NewAuthController(db *gorm.DB) AuthController {
	service := services.NewAuthService(db)
	controller := AuthController{
		authService: service,
		validate:    *vl.New(),
	}
	return controller
}

func (controller *AuthController) Login(c echo.Context) error {
	type payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	response := controller.authService.Login(payloadValidator.Email, payloadValidator.Password)
	return c.JSON(response.StatusCode, response)
}

func (controller *AuthController) Register(c echo.Context) error {
	type payload struct {
		NamaLengkap          string `json:"NamaLengkap" validate:"required"`
		Email                string `json:"Email" validate:"required,email"`
		Password             string `json:"Password" validate:"required"`
		PasswordConfirmation string `json:"PasswordConfirmation" validate:"required,eqfield=Password"`
		JenisKelamin         string `json:"JenisKelamin" validate:"required"`
		Umur                 int    `json:"Umur" validate:"required"`
		BeratBadan           int    `json:"BeratBadan" validate:"required"`
		TinggiBadan          int    `json:"TinggiBadan" validate:"required"`
		FrekuensiGym         int    `json:"FrekuensiGym" validate:"required"`
		TargetKalori         int    `json:"TargetKalori" validate:"required"`
		ReferalCode          string `json:"ReferalCode" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}
	var regisPayload utils.RegisterRequest = utils.RegisterRequest{
		Fullname:             payloadValidator.NamaLengkap,
		Email:                payloadValidator.Email,
		Password:             payloadValidator.Password,
		PasswordConfirmation: payloadValidator.PasswordConfirmation,
		JenisKelamin:         payloadValidator.JenisKelamin,
		Umur:                 payloadValidator.Umur,
		BeratBadan:           payloadValidator.BeratBadan,
		TinggiBadan:          payloadValidator.TinggiBadan,
		FrekuensiGym:         payloadValidator.FrekuensiGym,
		TargetKalori:         payloadValidator.TargetKalori,
	}

	response := controller.authService.Register(regisPayload)
	return c.JSON(response.StatusCode, response)
}

func (controller *AuthController) GetUser(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	fmt.Print("token", token)
	response := controller.authService.GetLoggedInUser(token)
	return c.JSON(response.StatusCode, response)
}
