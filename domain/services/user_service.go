package services

import (
	"fmt"
	"io"
	"kalorize-api/domain/repositories"
	"kalorize-api/utils"
	"os"
	"path/filepath"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	EditUser(token string, payload utils.UserRequest) utils.Response
	EditPassword(token string, payload utils.UserRequest) utils.Response
	EditPhoto(token string, payload utils.UploadedPhoto) utils.Response
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		userRepository: repositories.NewDBUserRepository(db),
	}
}

func (service *userService) EditUser(token string, payload utils.UserRequest) utils.Response {
	emailUser, err := utils.ParseData(token)
	if err != nil || emailUser == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}
	user, err := service.userRepository.GetUserByEmail(emailUser)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to get user",
			Data:       nil,
		}
	}
	user.Fullname = payload.Fullname
	user.Email = payload.Email
	user.BeratBadan = payload.BeratBadan
	user.TinggiBadan = payload.TinggiBadan
	user.Umur = payload.Umur
	user.FrekuensiGym = payload.FrekuensiGym
	user.TargetKalori = payload.TargetKalori
	err = service.userRepository.UpdateUser(user)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to edit user",
			Data:       nil,
		}
	}

	return utils.Response{
		StatusCode: 200,
		Messages:   "Success",
		Data:       user,
	}
}

func (service *userService) EditPassword(token string, payload utils.UserRequest) utils.Response {
	emailUser, err := utils.ParseData(token)
	if err != nil || emailUser == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}
	user, err := service.userRepository.GetUserByEmail(emailUser)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to get user",
			Data:       nil,
		}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Password hashing failed",
			Data:       nil,
		}
	}
	user.Password = string(hashedPassword)
	err = service.userRepository.UpdateUser(user)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to edit user",
			Data:       nil,
		}
	}

	return utils.Response{}
}

func (service *userService) EditPhoto(token string, payload utils.UploadedPhoto) utils.Response {
	emailUser, err := utils.ParseData(token)
	if err != nil || emailUser == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}
	user, err := service.userRepository.GetUserByEmail(emailUser)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to get user",
			Data:       nil,
		}
	}
	filename := payload.Handler.Filename
	if payload.Alias != "" {
		filename = fmt.Sprintf("%s%s", payload.Alias, filepath.Ext(payload.Handler.Filename))
	}
	dir, err := os.Getwd()
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to get current directory",
			Data:       nil,
		}
	}
	fileLocation := filepath.Join(dir, "storage", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to open file",
			Data:       nil,
		}
	}
	defer targetFile.Close()
	if _, err := io.Copy(targetFile, payload.File); err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to copy file",
			Data:       nil,
		}
	}
	user.Foto = payload.Alias + filepath.Ext(payload.Handler.Filename)
	return utils.Response{
		StatusCode: 200,
		Messages:   "Success",
		Data:       user.Foto,
	}
}
