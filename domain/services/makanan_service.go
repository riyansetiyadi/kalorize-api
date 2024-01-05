package services

import (
	"fmt"
	"kalorize-api/domain/models"
	"kalorize-api/domain/repositories"
	"kalorize-api/utils"

	"gorm.io/gorm"
)

type makananService struct {
	makananRepo repositories.MakananRepository
}

func (service *makananService) GetAllMakanan() utils.Response {
	var response utils.Response
	fmt.Print("masuk service")
	makanan, err := service.makananRepo.GetAllMakanan()
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

func (service *makananService) GetMakananById(id string) utils.Response {
	var response utils.Response
	makanan, err := service.makananRepo.GetMakananById(id)
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
}

func NewMakananService(db *gorm.DB) MakananService {
	return &makananService{makananRepo: repositories.NewDBMakananRepository(db)}
}
