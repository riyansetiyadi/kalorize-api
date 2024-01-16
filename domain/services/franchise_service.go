package services

import (
	"kalorize-api/domain/models"
	"kalorize-api/domain/repositories"
	"kalorize-api/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FranchiseService struct {
	franchiseRepo repositories.FranchiseRepository
}

func (service *FranchiseService) GetAllFranchise() utils.Response {
	var response utils.Response
	franchise, err := service.franchiseRepo.GetAllFranchise()
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Internal server error"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "success"
	response.Data = franchise
	return response
}

func (service *FranchiseService) GetFranchiseById(id string) utils.Response {
	var response utils.Response
	franchise, err := service.franchiseRepo.GetFranchiseById(id)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Internal server error"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "success"
	response.Data = franchise
	return response
}

func (service *FranchiseService) CreateFranchise(franchise models.Franchise) utils.Response {
	var response utils.Response
	err := service.franchiseRepo.CreateFranchise(franchise)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Internal server error"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "success"
	response.Data = franchise
	return response
}

func (service *FranchiseService) UpdateFranchise(franchise models.Franchise) utils.Response {
	var response utils.Response
	err := service.franchiseRepo.UpdateFranchise(franchise)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Internal server error"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "success"
	response.Data = franchise
	return response
}

func (service *FranchiseService) ConnectFranchiseToMakanan(idMakanan string, idFranchise uuid.UUID) utils.Response {
	var response utils.Response
	var franchiseMakanan models.FranchiseMakanan
	franchiseMakanan.IdFranchise = idFranchise
	franchiseMakanan.IdMakanan = idMakanan
	err := service.franchiseRepo.AddFranchiseMakanan(franchiseMakanan)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Internal server error"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "success"
	response.Data = franchiseMakanan
	return response
}

func NewFranchiseService(db *gorm.DB) *FranchiseService {
	repo := repositories.NewDBFranchiseRepository(db)
	return &FranchiseService{franchiseRepo: repo}
}
