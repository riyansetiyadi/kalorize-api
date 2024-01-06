package services

import (
	"kalorize-api/domain/repositories"
	"kalorize-api/utils"

	"gorm.io/gorm"
)

type GymService struct {
	gymRepo     repositories.GymRepository
	gymKode     repositories.KodeGymRepository
	gymUsedCode repositories.UsedCodeRepository
}

func (gymService *GymService) CheckGymCode(gymKode string) utils.Response {
	kodeGym, err := gymService.gymKode.GetIDFromKode(gymKode)
	if err != nil {
		return utils.Response{StatusCode: 500, Messages: err.Error()}
	}

	if kodeGym == 0 {
		return utils.Response{StatusCode: 404, Messages: "Kode Gym tidak ditemukan"}
	}

	if gymService.IsUsed(kodeGym) {
		return utils.Response{StatusCode: 400, Messages: "Kode Gym sudah digunakan"}
	}

	return utils.Response{StatusCode: 200, Messages: "Kode Gym valid"}
}

func (gymService *GymService) IsUsed(gymCode int) bool {
	usedCode, err := gymService.gymUsedCode.GetUsedCodeByIdCode(gymCode)
	if err != nil {
		return false
	}
	emptyUUID := 0
	if usedCode.IdKode == emptyUUID {
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
