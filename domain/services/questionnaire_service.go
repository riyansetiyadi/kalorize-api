package services

import (
	"kalorize-api/domain/repositories"
	"kalorize-api/utils"

	"gorm.io/gorm"
)

type questionnaireService struct {
	questionnaireRepo repositories.UserRepository
}
type QuestionnaireService interface {
	FillQuestionnaire(questionnaireRequest utils.UserRequest) utils.Response
}

func NewQuestionnaireService(db *gorm.DB) QuestionnaireService {
	return &questionnaireService{
		questionnaireRepo: repositories.NewDBUserRepository(db),
	}
}

func (service *questionnaireService) FillQuestionnaire(questionnaireRequest utils.UserRequest) utils.Response {
	var response utils.Response
	var user, err = service.questionnaireRepo.GetUserById(questionnaireRequest.IdUser)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "User is not found"
		response.Data = nil
		return response
	}
	user.Umur = questionnaireRequest.Umur
	user.BeratBadan = questionnaireRequest.BeratBadan
	user.TinggiBadan = questionnaireRequest.TinggiBadan
	user.JenisKelamin = questionnaireRequest.JenisKelamin
	user.FrekuensiGym = questionnaireRequest.FrekuensiGym
	user.TargetKalori = questionnaireRequest.TargetKalori
	err = service.questionnaireRepo.UpdateUser(user)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Failed to fill questionnaire"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "Success"
	response.Data = user
	return response
}
