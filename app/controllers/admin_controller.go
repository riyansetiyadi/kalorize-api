package controllers

import (
	"kalorize-api/app/services"
	"kalorize-api/utils"
	"net/http"
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
		NamaGym    string  `form:"namaGym" validate:"required"`
		AlamatGym  string  `form:"alamatGym" validate:"required"`
		Latitude   float64 `form:"latitude" validate:"required"`
		Longitude  float64 `form:"longitude" validate:"required"`
		LinkGoogle string  `form:"linkGoogle" validate:"required"`
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

	response := controller.adminService.RegisterGym(token, registGymPayload, photoRequest)
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
		Email        string `form:"email" validate:"required,email"`
		FullName     string `form:"fullname" validate:"required"`
		JenisKelamin int    `form:"jenis_kelamin" validate:"required"`
		NoTelepon    string `form:"no_telepon" validate:"required"`
		ReferalCode  string `form:"referal_code" validate:"required"`
		Umur         int    `form:"umur" validate:"required"`
		BeratBadan   int    `form:"berat_badan" validate:"required"`
		TinggiBadan  int    `form:"tinggi_badan" validate:"required"`
		FrekuensiGym int    `form:"frekuensi_gym" validate:"required"`
		TargetKalori int    `form:"target_kalori" validate:"required"`
		Password     string `form:"password" validate:"required"`
		Role         string `form:"role" validate:"required"`
	}
	payloadValidator := new(payload)
	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

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

	var registerUserPayload utils.UserRequest = utils.UserRequest{
		Email:        payloadValidator.Email,
		Password:     payloadValidator.Password,
		Role:         payloadValidator.Role,
		Fullname:     payloadValidator.FullName,
		JenisKelamin: payloadValidator.JenisKelamin,
		NoTelepon:    payloadValidator.NoTelepon,
		ReferalCode:  payloadValidator.ReferalCode,
		Umur:         payloadValidator.Umur,
		BeratBadan:   payloadValidator.BeratBadan,
		TinggiBadan:  payloadValidator.TinggiBadan,
		FrekuensiGym: payloadValidator.FrekuensiGym,
		TargetKalori: payloadValidator.TargetKalori,
	}
	response := controller.adminService.RegisterUser(token, registerUserPayload, photoRequest)
	return c.JSON(response.StatusCode, response)
}

func (controller *AdminController) GenerateGymToken(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	type payload struct {
		Uid uuid.UUID `json:"uid" validate:"required"`
	}
	payloadValidator := new(payload)
	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}
	response := controller.adminService.GenerateGymToken(token, payloadValidator.Uid)
	return c.JSON(response.StatusCode, response)
}

func (controller *AdminController) GetAllUser(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	response := controller.adminService.GetAllUser(token)
	return c.JSON(response.StatusCode, response)
}

func (controller *AdminController) GetUserById(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	id := c.Param("id")
	uuid := uuid.MustParse(id)
	response := controller.adminService.GetUserById(token, uuid)
	return c.JSON(response.StatusCode, response)
}

func (controller *AdminController) UpdateUser(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	id := c.Param("id")
	uuid := uuid.MustParse(id)
	type payload struct {
		FullName     string `json:"fullname"`
		Email        string `json:"email"`
		NoTelepon    string `json:"noTelepon"`
		Password     string `json:"password"`
		JenisKelamin int    `json:"jenisKelamin"`
		Umur         int    `json:"umur"`
		BeratBadan   int    `json:"beratBadan"`
		TinggiBadan  int    `json:"tinggiBadan"`
		FrekuensiGym int    `json:"frekuensiGym"`
		TargetKalori int    `json:"targetKalori"`
		ReferalCode  string `json:"referalCode"`
		Role         string `json:"role"`
	}
	payloadValidator := new(payload)
	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}
	var updateUserPayload utils.UserRequest = utils.UserRequest{
		Email:    payloadValidator.Email,
		Password: payloadValidator.Password,
		Role:     payloadValidator.Role,
	}
	response := controller.adminService.UpdateUser(token, uuid, updateUserPayload)
	return c.JSON(response.StatusCode, response)
}

func (controller *AdminController) DeleteUser(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	id := c.Param("id")
	uuid := uuid.MustParse(id)
	response := controller.adminService.DeleteUser(token, uuid)
	return c.JSON(response.StatusCode, response)
}
