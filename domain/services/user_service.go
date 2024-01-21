package services

import (
	"context"
	"fmt"
	"io"
	"kalorize-api/domain/models"
	"kalorize-api/domain/repositories"
	"kalorize-api/formatter"
	"kalorize-api/utils"
	"path/filepath"
	"reflect"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/option"
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
	userRepository     repositories.UserRepository
	historyRepository  repositories.HistoryRepository
	makananrRepository repositories.MakananRepository
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		userRepository:     repositories.NewDBUserRepository(db),
		historyRepository:  repositories.NewDBHistoryRepository(db),
		makananrRepository: repositories.NewDBMakananRepository(db),
	}
}

func (service *userService) CreateHistory(token string, historyPayload utils.HistoryRequest) utils.Response {
	emailUser, err := utils.ParseDataEmail(token)
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
		TotalProtein:  historyPayload.TotalProtein,
		TotalKalori:   historyPayload.TotalKalori,
		TanggalDibuat: time.Now(),
	}

	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to make history model",
			Data:       nil,
		}
	}
	err = service.historyRepository.CreateHistory(history)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to create history",
			Data:       nil,
		}
	}
	return utils.Response{
		StatusCode: 200,
		Messages:   "Success",
		Data:       history,
	}
}

func (service *userService) GetHistory(token string, date time.Time) utils.Response {
	emailUser, err := utils.ParseDataEmail(token)
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
	fmt.Print(history)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to get history",
			Data:       nil,
		}
	}
	breakfast, err := service.makananrRepository.GetMakananById(history.IdBreakfast)
	formattedBreakfast := formatter.FormatterMakananLuarIndo(breakfast)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to get breakfast",
			Data:       nil,
		}
	}
	lunch, err := service.makananrRepository.GetMakananById(history.IdLunch)
	formattedLunch := formatter.FormatterMakananLuarIndo(lunch)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to get lunch",
			Data:       nil,
		}
	}
	dinner, err := service.makananrRepository.GetMakananById(history.IdDinner)
	formattedDinner := formatter.FormatterMakananLuarIndo(dinner)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to get dinner",
			Data:       nil,
		}
	}
	var response utils.Response
	response.StatusCode = 200
	response.Messages = "Success"
	response.Data = map[string]interface{}{
		"breakfast":    formattedBreakfast,
		"lunch":        formattedLunch,
		"dinner":       formattedDinner,
		"totalKalori":  history.TotalKalori,
		"totalProtein": history.TotalProtein,
	}

	return response
}

func (service *userService) EditUser(token string, payload utils.UserRequest) utils.Response {
	emailUser, err := utils.ParseDataEmail(token)
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
	emailUser, err := utils.ParseDataEmail(token)
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
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
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
	emailUser, err := utils.ParseDataEmail(token)
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

	// Initialize Firebase app
	opt := option.WithCredentialsFile("config/credentials.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to initialize Firebase app",
			Data:       nil,
		}
	}

	// Initialize Firebase Storage client
	// Initialize Firebase Storage client
	client, err := app.Storage(context.Background())
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to initialize Firebase Storage client",
			Data:       nil,
		}
	}

	// Specify the path within the bucket where the file should be stored
	storagePath := fmt.Sprintf("images/%s", filename)

	// Open a new reader for the file
	reader := payload.File

	// Get the bucket handle from the client
	bucket, err := client.Bucket("kalorize-71324.appspot.com")
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to get bucket handle from the client",
			Data:       nil,
		}
	}
	// Initialize the writer for the file
	wc := bucket.Object(storagePath).NewWriter(context.Background())
	wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}

	// Upload the file to Firebase Storage
	if _, err := io.Copy(wc, reader); err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to upload file to Firebase Storage",
			Data:       nil,
		}
	}

	// Close the writer after copying
	if err := wc.Close(); err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to close Firebase Storage writer",
			Data:       nil,
		}
	}

	// Set user properties
	user.Foto = payload.Alias + filepath.Ext(payload.Handler.Filename)
	user.FotoUrl = fmt.Sprintf("https://storage.googleapis.com/kalorize-71324.appspot.com/%s", storagePath)

	// Update user in the database
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
