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

	if gymService.IsUsed(kodeGym.KodeGym) == (utils.Response{StatusCode: 200, Messages: "Kode Gym sudah digunakan"}) {
		return utils.Response{StatusCode: 400, Messages: "Kode Gym sudah digunakan"}
	}

	return utils.Response{StatusCode: 200, Messages: "Kode Gym valid"}
}

func (gymService *GymService) IsUsed(gymCode string) utils.Response {
	usedCode, err := gymService.gymUsedCode.GetUsedCodeByGymCode(gymCode)
	if err != nil {
		return utils.Response{StatusCode: 500, Messages: "Belum ada Gym terdaftar"}
	}
	if usedCode.KodeGym == "" {
		return utils.Response{StatusCode: 404, Messages: "Kode Gym tidak ditemukan"}
	}

	return utils.Response{StatusCode: 200, Messages: "Kode Gym sudah digunakan"}
}

func (gymService *GymService) GetAllGym() utils.Response {
	gym, _ := gymService.gymRepo.GetGym()
	if gym == nil {
		return utils.Response{StatusCode: 404, Messages: "Belum ada gym terdaftar"}
	}
	return utils.Response{StatusCode: 200, Messages: "Success", Data: gym}
}

func NewGymService(db *gorm.DB) *GymService {
	return &GymService{
		gymRepo:     repositories.NewDBGymRepository(db),
		gymKode:     repositories.NewDBKodeGymRepository(db),
		gymUsedCode: repositories.NewDBUsedCodeRepository(db),
	}
}
