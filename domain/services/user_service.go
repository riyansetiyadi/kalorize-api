package services

import (
	"fmt"
	"io"
	"kalorize-api/domain/models"
	"kalorize-api/domain/repositories"
	"kalorize-api/utils"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	GetHistory(token string, date time.Time) utils.Response
	CreateHistory(token string, historyPayload utils.HistoryRequest) utils.Response
	EditUser(token string, payload utils.UserRequest) utils.Response
	EditPassword(token string, payload utils.UserRequest, oldPassword string) utils.Response
	EditPhoto(token string, payload utils.UploadedPhoto) utils.Response
}

type userService struct {
	userRepository    repositories.UserRepository
	historyRepository repositories.HistoryRepository
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		userRepository:    repositories.NewDBUserRepository(db),
		historyRepository: repositories.NewDBHistoryRepository(db),
	}
}

func (service *userService) CreateHistory(token string, historyPayload utils.HistoryRequest) utils.Response {
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
	history := models.History{
		IdHistory:     uuid.New(),
		IdUser:        user.IdUser,
		IdBreakfast:   historyPayload.IdBreakfast,
		IdLunch:       historyPayload.IdLunch,
		IdDinner:      historyPayload.IdDinner,
		TanggalDibuat: time.Now(),
	}

	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to create history",
			Data:       nil,
		}
	}
	err = service.historyRepository.CreateHistory(history)
	return utils.Response{
		StatusCode: 200,
		Messages:   "Success",
		Data:       history,
	}
}

func (service *userService) GetHistory(token string, date time.Time) utils.Response {
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
	history, err := service.historyRepository.GetHistoryByIdUserAndDate(user.IdUser, date)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to get history",
			Data:       nil,
		}
	}
	return utils.Response{
		StatusCode: 200,
		Messages:   "Success",
		Data:       history,
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
	validateAndAssign(&user.Fullname, payload.Fullname)
	validateAndAssign(&user.Email, payload.Email)
	validateAndAssign(&user.NoTelepon, payload.NoTelepon)

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

func validateAndAssign(target interface{}, source interface{}) {
	if source != nil {
		targetValue := reflect.ValueOf(target)
		sourceValue := reflect.ValueOf(source)

		if targetValue.Kind() == reflect.Ptr && !targetValue.IsNil() {
			if sourceValue.Kind() == reflect.Ptr {
				if sourceValue.Elem().IsValid() {
					targetValue.Elem().Set(sourceValue.Elem())
				}
			} else {
				// If source is not a pointer, and not an empty string, directly set the value
				if sourceValue.Kind() != reflect.String || sourceValue.String() != "" {
					targetValue.Elem().Set(sourceValue)
				}
			}
		}
	}
}

func (service *userService) EditPassword(token string, payload utils.UserRequest, oldPassword string) utils.Response {
	emailUser, err := utils.ParseData(token)
	if err != nil || emailUser == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}
	user, err := service.userRepository.GetUserByEmail(emailUser)
	if err != nil && user.Email != emailUser {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to get user",
			Data:       nil,
		}
	}
	hashedOldPassword, err := bcrypt.GenerateFromPassword([]byte(oldPassword), bcrypt.DefaultCost)
	if user.Password != string(hashedOldPassword) {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Old password is wrong",
			Data:       nil,
		}
	}
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to validate old password",
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

	return utils.Response{
		StatusCode: 200,
		Messages:   "Success",
		Data:       user,
	}
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
	user.FotoUrl = "https://985e-36-71-83-68.ngrok-free.app/api/v1/storage/" + user.Foto
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
