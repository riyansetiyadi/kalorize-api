package controllers

import (
	"kalorize-api/domain/services"
	"kalorize-api/utils"
	"net/http"
	"strings"

	vl "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	userService services.UserService
	validate    vl.Validate
}

func NewUserController(db *gorm.DB) UserController {
	service := services.NewUserService(db)
	controller := UserController{
		userService: service,
		validate:    *vl.New(),
	}
	return controller
}

func (controller *UserController) EditUser(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	type payload struct {
		NamaUser     string `json:"namaUser" validate:"required"`
		EmailUser    string `json:"emailUser" validate:"required,email"`
		BeratBadan   int    `json:"beratBadan" validate:"required"`
		TinggiBadan  int    `json:"tinggiBadan" validate:"required"`
		Umur         int    `json:"umur" validate:"required"`
		FrekuensiGym int    `json:"frekuensiGym" validate:"required"`
		TargetKalori int    `json:"targetKalori" validate:"required"`
	}
	payloadValidator := new(payload)
	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}
	var editUserPayload utils.UserRequest = utils.UserRequest{
		Fullname:     payloadValidator.NamaUser,
		Email:        payloadValidator.EmailUser,
		BeratBadan:   payloadValidator.BeratBadan,
		TinggiBadan:  payloadValidator.TinggiBadan,
		Umur:         payloadValidator.Umur,
		FrekuensiGym: payloadValidator.FrekuensiGym,
		TargetKalori: payloadValidator.TargetKalori,
	}

	response := controller.userService.EditUser(token, editUserPayload)
	return c.JSON(response.StatusCode, response)
}

func (controller *UserController) EditPassword(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	type payload struct {
		PasswordUser             string `json:"passwordUser" validate:"required"`
		PasswordConfirmationUser string `json:"passwordConfirmationUser" validate:"required"`
	}
	payloadValidator := new(payload)
	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}
	var editPasswordPayload utils.UserRequest = utils.UserRequest{
		Password:             payloadValidator.PasswordUser,
		PasswordConfirmation: payloadValidator.PasswordConfirmationUser,
	}

	response := controller.userService.EditPassword(token, editPasswordPayload)
	return c.JSON(response.StatusCode, response)
}

func (controller *UserController) EditPhoto(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	// ParseMultipartForm with a maximum of 1024 bytes
	if err := c.Request().ParseMultipartForm(1024); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	alias := c.Request().FormValue("alias")
	uploadedFile, handler, err := c.Request().FormFile("file")
	if err != nil {
		return c.JSON(400, err.Error())
	}

	photoRequest := utils.UploadedPhoto{
		Alias:   alias,
		File:    uploadedFile,
		Handler: handler,
	}

	response := controller.userService.EditPhoto(token, photoRequest)
	return c.JSON(response.StatusCode, response)
}
