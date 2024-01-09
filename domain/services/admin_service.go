package services

import (
	"kalorize-api/domain/models"
	"kalorize-api/domain/repositories"
	"kalorize-api/utils"
	"strings"
	"time"

	"gorm.io/gorm"
)

type adminService struct {
	gymRepo       repositories.GymRepository
	userRepo      repositories.UserRepository
	gymKode       repositories.KodeGymRepository
	gymUsedCode   repositories.UsedCodeRepository
	makananRepo   repositories.MakananRepository
	franchiseRepo repositories.FranchiseRepository
}

func (service *adminService) RegisterGym(bearerToken string, registGymRequest utils.GymRequest) utils.Response {
	var response utils.Response
	adminEmail, err := utils.ParseData(bearerToken)
	if adminEmail == "" || err != nil {
		response.StatusCode = 401
		response.Messages = "Unauthorized"
		response.Data = nil
		return response
	}
	admin, err := service.userRepo.GetUserByEmail(adminEmail)
	if admin.Role != "admin" || err != nil {
		response.StatusCode = 401
		response.Messages = "Unauthorized"
		response.Data = nil
		return response
	}
	gym := models.Gym{
		NamaGym:      registGymRequest.NamaGym,
		AlamatGym:    registGymRequest.AlamatGym,
		EmailGym:     registGymRequest.EmailGym,
		PasswordGym:  registGymRequest.PasswordGym,
		NoTeleponGym: registGymRequest.NoTeleponGym,
	}

	err = service.gymRepo.CreateNewGym(gym)

	return utils.Response{}
}

func (service *adminService) RegisterFranchise(bearerToken string, registerFranchiseRequest utils.FranchiseRequest) utils.Response {
	var response utils.Response
	adminEmail, err := utils.ParseData(bearerToken)
	if adminEmail == "" || err != nil {
		response.StatusCode = 401
		response.Messages = "Unauthorized"
		response.Data = nil
		return response
	}
	admin, err := service.userRepo.GetUserByEmail(adminEmail)
	if admin.Role != "admin" || err != nil {
		response.StatusCode = 401
		response.Messages = "Unauthorized"
		response.Data = nil
		return response
	}
	franchise := models.Franchise{
		NamaFranchise:      registerFranchiseRequest.NamaFranchise,
		AlamatFranchise:    registerFranchiseRequest.AlamatFranchise,
		EmailFranchise:     registerFranchiseRequest.EmailFranchise,
		PasswordFranchise:  registerFranchiseRequest.PasswordFranchise,
		NoTeleponFranchise: registerFranchiseRequest.NoTeleponFranchise,
	}
	err = service.franchiseRepo.CreateFranchise(franchise)
	return utils.Response{}
}

func (service *adminService) RegisterMakanan(bearerToken string, registMakananRequest utils.MakananRequest) utils.Response {
	var response utils.Response
	adminEmail, err := utils.ParseData(bearerToken)
	if adminEmail == "" || err != nil {
		response.StatusCode = 401
		response.Messages = "Unauthorized"
		response.Data = nil
		return response
	}
	admin, err := service.userRepo.GetUserByEmail(adminEmail)
	if admin.Role != "admin" || err != nil {
		response.StatusCode = 401
		response.Messages = "Unauthorized"
		response.Data = nil
		return response
	}
	makanan := models.Makanan{
		Nama:        registMakananRequest.Nama,
		Kalori:      registMakananRequest.Kalori,
		Protein:     registMakananRequest.Protein,
		Jenis:       registMakananRequest.Jenis,
		Bahan:       strings.Join(registMakananRequest.Bahan, ", "),
		CookingStep: strings.Join(registMakananRequest.CookingStep, "., "),
		CreatedAt:   models.TimeWrapper{time.Now()},
		UpdatedAt:   models.TimeWrapper{time.Now()},
	}
	err = service.makananRepo.CreateMakanan(makanan)
	return utils.Response{}
}

type AdminService interface {
	RegisterGym(bearerToken string, registGymRequest utils.GymRequest) utils.Response
	RegisterFranchise(bearerToken string, registFranchiseRequest utils.FranchiseRequest) utils.Response
	RegisterMakanan(bearerToken string, registMakananRequest utils.MakananRequest) utils.Response
}

func NewAdminService(db *gorm.DB) AdminService {
	return &adminService{
		gymRepo:       repositories.NewDBGymRepository(db),
		gymKode:       repositories.NewDBKodeGymRepository(db),
		gymUsedCode:   repositories.NewDBUsedCodeRepository(db),
		makananRepo:   repositories.NewDBMakananRepository(db),
		franchiseRepo: repositories.NewDBFranchiseRepository(db),
	}
}
