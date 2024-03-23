package services

import (
	"fmt"
	"kalorize-api/app/models"
	"kalorize-api/app/repositories"
	"kalorize-api/utils"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
	authRepo     repositories.UserRepository
	usedCodeRepo repositories.UsedCodeRepository
	tokenRepo    repositories.TokenRepository
	gymRepo      repositories.GymRepository
}

func (service *authService) Login(email, password string) utils.Response {
	var response utils.Response
	if email == "" || password == "" {
		response.StatusCode = 400
		response.Messages = "email dan password tidak boleh kosong"
		response.Data = nil
		return response
	}

	if !utils.IsEmailValid(email) {
		response.StatusCode = 400
		response.Messages = "Email kamu tidak valid"
		response.Data = nil
		return response
	}

	user, err := service.authRepo.GetUserByEmail(email)
	if err != nil {
		response.StatusCode = 401
		response.Messages = "Email kamu belum terdaftar"
		response.Data = nil
		return response
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		fmt.Print(!utils.CheckPasswordHash(password, user.Password))
		response.StatusCode = 401
		response.Messages = "Password kamu salah"
		response.Data = nil
		return response
	}
	AccessToken, err := utils.GenerateJWTAccessToken(user.IdUser, user.Fullname, user.Email, "kalorize")
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token generation failed"
		response.Data = nil
		return response
	}
	refreshToken, err := utils.GenerateJWTRefreshToken(user.IdUser, user.Fullname, user.Email, "kalorize")
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token generation failed"
		response.Data = nil
		return response
	}
	token := models.Token{
		IdToken:      uuid.New(),
		AccessToken:  AccessToken,
		RefreshToken: refreshToken,
		UserId:       user.IdUser,
	}
	err = service.tokenRepo.CreateNewToken(token)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token creation failed"
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "success"
	response.Data = map[string]interface{}{
		"accessToken":  AccessToken,
		"refreshToken": refreshToken,
		"role":         user.Role,
		"userId":       user.IdUser,
	}
	return response
}

func (service *authService) Register(registerRequest utils.UserRequest, gymKode string) utils.Response {
	var response utils.Response
	if registerRequest.Fullname == "" || registerRequest.Email == "" || registerRequest.Password == "" || registerRequest.PasswordConfirmation == "" {
		response.StatusCode = 400
		response.Messages = "Semua field harus diisi"
		response.Data = nil
		return response
	}
	user, err := service.authRepo.GetUserByEmail(registerRequest.Email)
	if err == nil {
		response.StatusCode = 400
		response.Messages = "Email sudah terdaftar"
		response.Data = nil
		return response
	}

	if !utils.IsEmailValid(registerRequest.Email) {
		response.StatusCode = 400
		response.Messages = "Email kamu tidak valid"
		response.Data = nil
		return response
	}

	// if service.authRepo.FindReferalCodeIfExist(registerRequest.ReferalCode) == true {
	// 	adain special service untuk user
	// }

	if registerRequest.Password != registerRequest.PasswordConfirmation {
		response.StatusCode = 400
		response.Messages = "Password dan konfirmasi password tidak sama"
		response.Data = nil
		return response
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Password hashing failed"
		response.Data = nil
		return response
	}
	userId := uuid.New()
	accessToken, err := utils.GenerateJWTAccessToken(userId, registerRequest.Fullname, registerRequest.Email, "kalorize")
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token generation failed"
		response.Data = nil
		return response
	}

	refreshtoken, err := utils.GenerateJWTRefreshToken(userId, registerRequest.Fullname, registerRequest.Email, "kalorize")
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token generation failed"
		response.Data = nil
		return response
	}

	user = models.User{
		IdUser:      userId,
		Fullname:    registerRequest.Fullname,
		Email:       registerRequest.Email,
		Umur:        registerRequest.Umur,
		ReferalCode: utils.GenerateReferalCode(registerRequest.Fullname),
		Password:    string(hashedPassword),
		Role:        registerRequest.Role,
	}
	uuidString := "10bedc93-46f9-4111-87ec-c9ad948aff81"
	parsedUUID, err := uuid.Parse(uuidString)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gym tidak ditemukan"
		response.Data = nil
		return response
	}
	gym, err := service.gymRepo.GetGymById(parsedUUID)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gym tidak ditemukan"
		response.Data = nil
		return response
	}
	usedCode := models.UsedCode{
		IdGym:   gym.IdGym,
		KodeGym: gymKode,
		IdUser:  user.IdUser,
	}
	err = service.usedCodeRepo.CreateNewUsedCode(usedCode)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Used code creation failed"
		response.Data = user.IdUser
		return response
	}
	err = service.authRepo.CreateNewUser(user)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "User creation failed"
		response.Data = user.IdUser
		return response
	}
	response.StatusCode = 200
	response.Messages = "success"
	response.Data = map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshtoken,
		"role":         user.Role,
		"userId":       user.IdUser,
	}
	return response
}

func (service *authService) GetLoggedInUser(bearerToken string) utils.Response {
	var response utils.Response
	var firstname, lastname string
	id, err := utils.ParseDataId(bearerToken)
	if id != uuid.Nil && err == nil {
		user, err := service.authRepo.GetUserById(id)
		if err != nil {
			response.StatusCode = 500
			response.Messages = "User tidak ditemukan"
			response.Data = nil
			return response
		}

		response.StatusCode = 200
		response.Messages = "success"
		names := strings.Split(user.Fullname, " ")
		if len(names) == 1 {
			firstname = names[0]
			lastname = names[0]
		} else {
			firstname = names[0]
			lastname = names[len(names)-1]
		}
		KodeGym, err := service.usedCodeRepo.GetusedCodeByIdUser(user.IdUser)
		if err != nil {
			response.StatusCode = 500
			response.Messages = "Kode gym tidak ditemukan"
			response.Data = nil
			return response
		}
		Gym, err := service.gymRepo.GetGymById(KodeGym.IdGym)
		if err != nil {
			response.StatusCode = 500
			response.Messages = "Gym tidak ditemukan"
			response.Data = nil
			return response
		}
		response.Data = map[string]interface{}{
			"idUser":       user.IdUser,
			"firstName":    firstname,
			"lastName":     lastname,
			"email":        user.Email,
			"jenisKelamin": user.JenisKelamin,
			"frekuensiGym": user.FrekuensiGym,
			"targetKalori": user.TargetKalori,
			"tinggiBadan":  user.TinggiBadan,
			"umur":         user.Umur,
			"beratBadan":   user.BeratBadan,
			"role":         user.Role,
			"foto":         user.FotoUrl,
			"noTelepon":    user.NoTelepon,
			"KodeGym":      KodeGym.KodeGym,
			"Gym":          Gym.NamaGym,
		}
		return response
	} else {
		response.StatusCode = 401
		response.Messages = "Invalid token"
		response.Data = nil
		return response
	}
}

func (service *authService) Logout(bearerToken string) utils.Response {
	var response utils.Response
	_, err := utils.ParseDataEmail(bearerToken)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token parsing failed"
		response.Data = nil
		return response
	}
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token deletion failed"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "success"
	response.Data = nil
	return response
}

type AuthService interface {
	Login(username, password string) utils.Response
	Register(requestRegister utils.UserRequest, gymKode string) utils.Response
	GetLoggedInUser(bearerToken string) utils.Response
	Logout(bearerToken string) utils.Response
}

func NewAuthService(db *gorm.DB) AuthService {
	return &authService{
		authRepo:     repositories.NewDBUserRepository(db),
		usedCodeRepo: repositories.NewDBUsedCodeRepository(db),
		tokenRepo:    repositories.NewDBTokenRepository(db),
		gymRepo:      repositories.NewDBGymRepository(db),
	}
}
