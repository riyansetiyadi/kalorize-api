package repositories

import (
	"kalorize-api/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DbToken struct {
	Conn *gorm.DB
}

func (db *DbToken) GetToken() ([]models.Token, error) {
	var token []models.Token
	err := db.Conn.Find(&token).Error
	return token, err
}

func (db *DbToken) CreateNewToken(token models.Token) error {
	return db.Conn.Create(&token).Error
}

func (db *DbToken) UpdateToken(token models.Token) error {
	return db.Conn.Save(&token).Error
}

func (db *DbToken) DeleteToken(idToken uuid.UUID) error {
	return db.Conn.Delete(&models.Token{}, idToken).Error
}

type TokenRepository interface {
	GetToken() ([]models.Token, error)
	CreateNewToken(token models.Token) error
	UpdateToken(models.Token) error
	DeleteToken(idToken uuid.UUID) error
}

func NewDBTokenRepository(conn *gorm.DB) *DbToken {
	return &DbToken{Conn: conn}
}
