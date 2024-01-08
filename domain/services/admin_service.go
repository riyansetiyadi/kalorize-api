package services

import (
	"kalorize-api/domain/repositories"
	"kalorize-api/utils"
)

type AdminService struct {
	gymRepo     repositories.GymRepository
	gymKode     repositories.KodeGymRepository
	gymUsedCode repositories.UsedCodeRepository
}

func (adminService *AdminService) CreateNewGym() utils.Response {
	// var response utils.Response
	// err := adminService.gymRepo.CreateNewGym()
	// if err != nil {
	// 	response.StatusCode = 500
	// 	response.Messages = "Failed to create new gym"
	// 	response.Data = nil
	// 	return response
	// }
	// response.StatusCode = 200
	// response.Messages = "Success"
	// response.Data = nil
	// return response
	return utils.Response{}
}
