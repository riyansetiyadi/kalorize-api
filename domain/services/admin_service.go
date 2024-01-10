package services

import (
	"kalorize-api/domain/models"
	"kalorize-api/domain/repositories"
	"kalorize-api/utils"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func NewAdminService(db *gorm.DB) AdminService {
	return &adminService{
		userRepo:      repositories.NewDBUserRepository(db),
		gymRepo:       repositories.NewDBGymRepository(db),
		gymKode:       repositories.NewDBKodeGymRepository(db),
		gymUsedCode:   repositories.NewDBUsedCodeRepository(db),
		makananRepo:   repositories.NewDBMakananRepository(db),
		franchiseRepo: repositories.NewDBFranchiseRepository(db),
	}
}

func (service *adminService) RegisterGym(token string, registGymRequest utils.GymRequest) utils.Response {
	var response utils.Response
	adminEmail, err := utils.ParseData(token)
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registGymRequest.PasswordGym), bcrypt.DefaultCost)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Password hashing failed"
		response.Data = nil
		return response
	}

	gym := models.Gym{
		IdGym:        uuid.New(),
		NamaGym:      registGymRequest.NamaGym,
		AlamatGym:    registGymRequest.AlamatGym,
		EmailGym:     registGymRequest.EmailGym,
		PasswordGym:  string(hashedPassword),
		NoTeleponGym: registGymRequest.NoTeleponGym,
	}
	err = service.gymRepo.CreateNewGym(gym)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Failed to create gym"
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "Success"
	response.Data = gym
	return response
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerFranchiseRequest.PasswordFranchise), bcrypt.DefaultCost)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Password hashing failed"
		response.Data = nil
		return response
	}

	franchise := models.Franchise{
		IdFranchise:        uuid.New(),
		NamaFranchise:      registerFranchiseRequest.NamaFranchise,
		EmailFranchise:     registerFranchiseRequest.EmailFranchise,
		LongitudeFranchise: registerFranchiseRequest.LongitudeFranchise,
		LatitudeFranchise:  registerFranchiseRequest.LatitudeFranchise,
		LokasiFranchise:    registerFranchiseRequest.LokasiFranchise,
		FotoFranchise:      registerFranchiseRequest.FotoFranchise,
		PasswordFranchise:  string(hashedPassword),
		NoTeleponFranchise: registerFranchiseRequest.NoTeleponFranchise,
	}
	err = service.franchiseRepo.CreateFranchise(franchise)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Failed to create franchise"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "Success"
	response.Data = franchise
	return response
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
		CreatedAt:   models.TimeWrapper{Time: time.Now()},
		UpdatedAt:   models.TimeWrapper{Time: time.Now()},
	}
	err = service.makananRepo.CreateMakanan(makanan)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Failed to create makanan"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "Success"
	response.Data = makanan
	return response
}

type AdminService interface {
	RegisterGym(bearerToken string, registGymRequest utils.GymRequest) utils.Response
	RegisterFranchise(bearerToken string, registFranchiseRequest utils.FranchiseRequest) utils.Response
	RegisterMakanan(bearerToken string, registMakananRequest utils.MakananRequest) utils.Response
}
