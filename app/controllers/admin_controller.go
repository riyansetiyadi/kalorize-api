package controllers

import (
	"kalorize-api/app/services"
	"kalorize-api/utils"
	"strings"

	vl "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AdminController struct {
	adminService services.AdminService
	validate     vl.Validate
}

func NewAdminController(db *gorm.DB) AdminController {
	service := services.NewAdminService(db)
	controller := AdminController{
		adminService: service,
		validate:     *vl.New(),
	}
	return controller
}

func (controller *AdminController) RegisterGym(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	type payload struct {
		NamaGym    string  `json:"namaGym" validate:"required"`
		AlamatGym  string  `json:"alamatGym" validate:"required"`
		Latitude   float64 `json:"latitude" validate:"required"`
		Longitude  float64 `json:"longitude" validate:"required"`
		LinkGoogle string  `json:"linkGoogle" validate:"required"`
	}
	payloadValidator := new(payload)
	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}
	var registGymPayload utils.GymRequest = utils.GymRequest{
		NamaGym:    payloadValidator.NamaGym,
		AlamatGym:  payloadValidator.AlamatGym,
		Latitude:   payloadValidator.Latitude,
		Longitude:  payloadValidator.Longitude,
		LinkGoogle: payloadValidator.LinkGoogle,
	}

	response := controller.adminService.RegisterGym(token, registGymPayload)
	return c.JSON(response.StatusCode, response)
}

func (controller *AdminController) RegisterFranchise(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	type payload struct {
		NamaFranchise      string  `json:"namaFranchise" validate:"required"`
		LongitudeFranchise float64 `json:"longitudeFranchise" validate:"required"`
		LatitudeFranchise  float64 `json:"latitudeFranchise" validate:"required"`
		EmailFranchise     string  `json:"emailFranchise" validate:"required,email"`
		PasswordFranchise  string  `json:"passwordFranchise" validate:"required"`
		NoTeleponFranchise string  `json:"noTeleponFranchise" validate:"required"`
		FotoFranchise      string  `json:"fotoFranchise" validate:"required"`
		LokasiFranchise    string  `json:"lokasiFranchise" validate:"required"`
	}
	payloadValidator := new(payload)
	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}
	var registerFranchisePayload utils.FranchiseRequest = utils.FranchiseRequest{
		NamaFranchise:      payloadValidator.NamaFranchise,
		EmailFranchise:     payloadValidator.EmailFranchise,
		PasswordFranchise:  payloadValidator.PasswordFranchise,
		NoTeleponFranchise: payloadValidator.NoTeleponFranchise,
		LongitudeFranchise: payloadValidator.LongitudeFranchise,
		LatitudeFranchise:  payloadValidator.LatitudeFranchise,
		FotoFranchise:      payloadValidator.FotoFranchise,
		LokasiFranchise:    payloadValidator.LokasiFranchise,
	}
	response := controller.adminService.RegisterFranchise(token, registerFranchisePayload)
	return c.JSON(response.StatusCode, response)
}

func (controller *AdminController) RegisterMakanan(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	type payload struct {
		NamaMakanan   string   `json:"namaMakanan" validate:"required"`
		Kalori        int      `json:"kalori" validate:"required"`
		Protein       int      `json:"protein" validate:"required"`
		Bahan         []string `json:"bahan" validate:"required"`
		ListFranchise []string `json:"listFranchise" validate:"required"`
		CookingStep   []string `json:"cookingStep" validate:"required"`
	}
	payloadValidator := new(payload)
	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}
	var registerMakananPayload utils.MakananRequest = utils.MakananRequest{
		Nama:        payloadValidator.NamaMakanan,
		Kalori:      payloadValidator.Kalori,
		Protein:     payloadValidator.Protein,
		Bahan:       payloadValidator.Bahan,
		CookingStep: payloadValidator.CookingStep,
	}
	response := controller.adminService.RegisterMakanan(token, registerMakananPayload)
	return c.JSON(response.StatusCode, response)
}

func (controller *AdminController) RegisterUser(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	type payload struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required"`
	}
	payloadValidator := new(payload)
	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}
	var registerUserPayload utils.UserRequest = utils.UserRequest{
		Email:    payloadValidator.Email,
		Password: payloadValidator.Password,
		Role:     payloadValidator.Role,
	}
	response := controller.adminService.RegisterUser(token, registerUserPayload)
	return c.JSON(response.StatusCode, response)
}

func (controller *AdminController) GenerateGymToken(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	uuidString := "10bedc93-46f9-4111-87ec-c9ad948aff81"
	parsedUUID, err := uuid.Parse(uuidString)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	response := controller.adminService.GenerateGymToken(token, parsedUUID)
	return c.JSON(response.StatusCode, response)
}
