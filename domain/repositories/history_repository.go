package repositories

import (
	"kalorize-api/domain/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbHistory struct {
	Conn *gorm.DB
}

func (db *dbHistory) GetAllHistory() ([]models.History, error) {
	var history []models.History
	err := db.Conn.Find(&history).Error
	return history, err
}

func (db *dbHistory) GetHistoryById(id string) (models.History, error) {
	var history models.History
	err := db.Conn.Where("id_history = ?", id).First(&history).Error
	return history, err
}

func (db *dbHistory) CreateHistory(history models.History) error {
	return db.Conn.Create(&history).Error
}

func (db *dbHistory) UpdateHistory(history models.History) error {
	return db.Conn.Save(&history).Error
}

func (db *dbHistory) DeleteHistory(id string) error {
	return db.Conn.Delete(&models.History{}, id).Error
}

func (db *dbHistory) GetHistoryByIdUser(id uuid.UUID) ([]models.History, error) {
	var history []models.History
	err := db.Conn.Where("id_user = ?", id).Find(&history).Error
	return history, err
}

func (db *dbHistory) GetHistoryByIdUserAndDate(id uuid.UUID, date time.Time) ([]models.History, error) {
	var history []models.History
	err := db.Conn.Where("id_user = ? AND date = ?", id, date).Find(&history).Error
	return history, err
}

type HistoryRepository interface {
	GetAllHistory() ([]models.History, error)
	GetHistoryById(id string) (models.History, error)
	CreateHistory(history models.History) error
	UpdateHistory(history models.History) error
	DeleteHistory(id string) error
	GetHistoryByIdUser(id uuid.UUID) ([]models.History, error)
	GetHistoryByIdUserAndDate(id uuid.UUID, date time.Time) ([]models.History, error)
}

func NewDBHistoryRepository(conn *gorm.DB) *dbHistory {
	return &dbHistory{Conn: conn}
}
