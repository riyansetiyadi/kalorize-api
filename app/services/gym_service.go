package services

import (
	"kalorize-api/app/repositories"
	"kalorize-api/utils"

	"gorm.io/gorm"
)

type GymService struct {
	gymRepo     repositories.GymRepository
	gymKode     repositories.KodeGymRepository
	gymUsedCode repositories.UsedCodeRepository
}

func (gymService *GymService) CheckGymCode(gymKode string) utils.Response {
	kodeGym, err := gymService.gymKode.GetKodeGymByKode(gymKode)
	if err != nil {
		return utils.Response{StatusCode: 500, Messages: err.Error()}
	}

	if kodeGym.KodeGym == "" {
		return utils.Response{StatusCode: 404, Messages: "Kode Gym tidak ditemukan"}
	}

	if gymService.IsUsed(kodeGym.KodeGym) {
		return utils.Response{StatusCode: 400, Messages: "Kode Gym sudah digunakan"}
	}

	return utils.Response{StatusCode: 200, Messages: "Kode Gym valid"}
}

func (gymService *GymService) IsUsed(gymCode string) bool {
	usedCode, err := gymService.gymUsedCode.GetUsedCodeByGymCode(gymCode)
	if err != nil {
		return false
	}
	if usedCode.KodeGym == "" {
		return false
	}

	return true
}

func NewGymService(db *gorm.DB) *GymService {
	return &GymService{
		gymRepo:     repositories.NewDBGymRepository(db),
		gymKode:     repositories.NewDBKodeGymRepository(db),
		gymUsedCode: repositories.NewDBUsedCodeRepository(db),
	}
}
