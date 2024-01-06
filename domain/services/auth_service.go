package services

import (
	"fmt"
	"kalorize-api/domain/models"
	"kalorize-api/domain/repositories"
	"kalorize-api/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
	authRepo repositories.AuthRepository
}

func (service *authService) Login(email, password string) utils.Response {
	var response utils.Response
	if email == "" || password == "" {
		response.StatusCode = 400
		response.Messages = "Username dan password tidak boleh kosong"
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
	token, err := utils.GenerateJWTToken(user.Fullname, user.Email, "kalorize")
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token generation failed"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "success"

	response.Data = map[string]interface{}{
		"token": token,
	}
	return response

}

func (service *authService) Register(registerRequest utils.RegisterRequest) utils.Response {
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
	token, err := utils.GenerateJWTToken(registerRequest.Fullname, registerRequest.Email, "kalorize")
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token generation failed"
		response.Data = nil
		return response
	}

	response.Data = map[string]interface{}{
		"token": token,
	}
	uuid := uuid.New()
	user = models.User{
		IdUser:       uuid,
		Fullname:     registerRequest.Fullname,
		Email:        registerRequest.Email,
		Umur:         registerRequest.Umur,
		BeratBadan:   registerRequest.BeratBadan,
		TinggiBadan:  registerRequest.TinggiBadan,
		FrekuensiGym: registerRequest.FrekuensiGym,
		TargetKalori: registerRequest.TargetKalori,
		JenisKelamin: registerRequest.JenisKelamin,
		ReferalCode:  utils.GenerateReferalCode(registerRequest.Fullname),
		Password:     string(hashedPassword),
	}
	err = service.authRepo.CreateNewUser(user)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "User creation failed"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "success"
	return response
}

func (service *authService) GetLoggedInUser(bearerToken string) utils.Response {
	var response utils.Response
	fmt.Println(bearerToken)
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("kalorize"), nil
	})

	if err != nil {
		response.StatusCode = 401
		response.Messages = "Invalid token"
		response.Data = nil
		return response
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Print(claims)
		emailClaim := claims["Email"]
		fmt.Println(emailClaim)
		if emailClaim == nil {
			response.StatusCode = 401
			response.Messages = "Invalid token"
			response.Data = nil
			return response
		}

		email := emailClaim.(string)
		user, err := service.authRepo.GetUserByEmail(email)
		fmt.Println(user)
		if err != nil {
			response.StatusCode = 500
			response.Messages = "User tidak ditemukan"
			response.Data = nil
			return response
		}

		response.StatusCode = 200
		response.Messages = "success"
		names := strings.Split(user.Fullname, " ")
		firstname := names[0]
		lastname := names[1]

		response.Data = map[string]interface{}{
			"firstName":    firstname,
			"lastName":     lastname,
			"email":        user.Email,
			"jenisKelamin": user.JenisKelamin,
			"frekuensiGym": user.FrekuensiGym,
			"targetKalori": user.TargetKalori,
			"tinggiBadan":  user.TinggiBadan,
			"umur":         user.Umur,
			"beratBadan":   user.BeratBadan,
		}
		return response
	} else {
		response.StatusCode = 401
		response.Messages = "Invalid token"
		response.Data = nil
		return response
	}
}

type AuthService interface {
	Login(username, password string) utils.Response
	Register(requestRegister utils.RegisterRequest) utils.Response
	GetLoggedInUser(bearerToken string) utils.Response
}

func NewAuthService(db *gorm.DB) AuthService {
	return &authService{authRepo: repositories.NewDBAuthRepository(db)}
}
