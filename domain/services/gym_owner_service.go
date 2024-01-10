package services

import (
	"kalorize-api/domain/models"
	"kalorize-api/domain/repositories"
	"kalorize-api/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type gymOwnerService struct {
	gymRepo     repositories.GymRepository
	gymKode     repositories.KodeGymRepository
	gymUsedCode repositories.UsedCodeRepository
}

func (gymOwner *gymOwnerService) GenerateKodeGym(idGym uuid.UUID) utils.Response {
	var response utils.Response
	Gym, err := gymOwner.gymRepo.GetGymById(idGym)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Failed to get gym"
		response.Data = nil
		return response
	}
	var kodeGym = models.KodeGym{
		IdKodeGym:   uuid.New(),
		KodeGym:     utils.GenerateKodeGym(Gym.NamaGym),
		IdGym:       idGym,
		ExpiredTime: utils.GetExpiredTime(),
	}

	err = gymOwner.gymKode.CreateNewKodeGym(kodeGym)

	if err == nil {
		response.StatusCode = 500
		response.Messages = "Failed to generate kode gym"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "Success"
	response.Data = kodeGym
	return response
}

type GymOwnerService interface {
	GenerateKodeGym(idGym uuid.UUID) utils.Response
}

func NewGymOwnerService(db *gorm.DB) GymOwnerService {
	return &gymOwnerService{
		gymRepo:     repositories.NewDBGymRepository(db),
		gymKode:     repositories.NewDBKodeGymRepository(db),
		gymUsedCode: repositories.NewDBUsedCodeRepository(db),
	}
}
