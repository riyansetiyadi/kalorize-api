package services

import (
	"encoding/csv"
	"kalorize-api/domain/models"
	"kalorize-api/domain/repositories"
	"kalorize-api/formatter"
	"kalorize-api/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type makananService struct {
	makananRepo repositories.MakananRepository
}

func (service *makananService) GetAllMakanan() utils.Response {
	var response utils.Response
	makanan, err := service.makananRepo.GetAllMakanan()
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Internal server error"
		response.Data = nil
		return response
	}
	var formattedMakanan []formatter.MakananFormat
	for i := range makanan {
		if makanan[i].IdMakanan > "0" && makanan[i].IdMakanan < "201" {
			formattedMakanan = append(formattedMakanan, formatter.FormatterMakananIndo(makanan[i]))
		} else {
			formattedMakanan = append(formattedMakanan, formatter.FormatterMakananLuarIndo(makanan[i]))
		}
	}
	response.StatusCode = 200
	response.Messages = "success"
	response.Data = formattedMakanan
	return response
}

func (service *makananService) GetMakananCSV(c echo.Context) utils.Response {
	// response is .csv file generator
	wr := csv.NewWriter(c.Response())
	var response utils.Response
	makanan, err := service.makananRepo.GetAllMakanan()
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Internal server error"
		response.Data = nil
		return response
	}
	formattedMultiMakanan := formatter.FormatterMakananToMultiDimentionalArray(makanan)
	wr.WriteAll(formattedMultiMakanan)

	response.StatusCode = 200
	response.Messages = "success"
	return response
}

func (service *makananService) GetMakananById(id string) utils.Response {
	var response utils.Response
	makanan, err := service.makananRepo.GetMakananById(id)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Internal server error"
		response.Data = nil
		return response
	}
	var formattedMakanan formatter.MakananFormat
	if id > "0" && id < "201" {
		formattedMakanan = formatter.FormatterMakananIndo(makanan)
	} else {
		formattedMakanan = formatter.FormatterMakananLuarIndo(makanan)

	}

	response.StatusCode = 200
	response.Messages = "success"
	response.Data = formattedMakanan
	return response
}

func (service *makananService) CreateMakanan(makanan models.Makanan) utils.Response {
	var response utils.Response
	err := service.makananRepo.CreateMakanan(makanan)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Internal server error"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "success"
	response.Data = makanan
	return response
}

type MakananService interface {
	GetAllMakanan() utils.Response
	GetMakananById(id string) utils.Response
	CreateMakanan(makanan models.Makanan) utils.Response
	GetMakananCSV(c echo.Context) utils.Response
}

func NewMakananService(db *gorm.DB) MakananService {
	return &makananService{makananRepo: repositories.NewDBMakananRepository(db)}
}
