package services

import (
	"fmt"
	"kalorize-api/domain/auth/models"
	"kalorize-api/domain/auth/repositories"
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
	if !CheckPasswordHash(password, user.Password) {
		fmt.Print(!CheckPasswordHash(password, user.Password))
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

func (service *authService) Register(fullname, email, password, passwordConfirmation string) utils.Response {
	var response utils.Response
	if fullname == "" || email == "" || password == "" || passwordConfirmation == "" {
		response.StatusCode = 400
		response.Messages = "Semua field harus diisi"
		response.Data = nil
		return response
	}
	user, err := service.authRepo.GetUserByEmail(email)
	if err == nil {
		response.StatusCode = 400
		response.Messages = "Email sudah terdaftar"
		response.Data = nil
		return response
	}

	//check wether the email is valid or not
	if !utils.IsEmailValid(email) {
		response.StatusCode = 400
		response.Messages = "Email kamu tidak valid"
		response.Data = nil
		return response
	}

	if password != passwordConfirmation {
		response.StatusCode = 400
		response.Messages = "Password dan konfirmasi password tidak sama"
		response.Data = nil
		return response
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Password hashing failed"
		response.Data = nil
		return response
	}
	token, err := utils.GenerateJWTToken(fullname, email, "kalorize")
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
		IdUser:   uuid,
		Fullname: fullname,
		Email:    email,
		Password: string(hashedPassword),
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

		response.Data = map[string]string{
			"firstName": firstname,
			"lastName":  lastname,
		}
		return response
	} else {
		response.StatusCode = 401
		response.Messages = "Invalid token"
		response.Data = nil
		return response
	}
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type AuthService interface {
	Login(username, password string) utils.Response
	Register(fullname, email, password, passwordConfirmation string) utils.Response
	GetLoggedInUser(bearerToken string) utils.Response
}

func NewAuthService(db *gorm.DB) AuthService {
	return &authService{authRepo: repositories.NewDBAuthRepository(db)}
}
