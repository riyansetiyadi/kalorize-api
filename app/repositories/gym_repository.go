package repositories

import (
	"kalorize-api/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbGym struct {
	Conn *gorm.DB
}

func (db *dbGym) GetGym() ([]models.Gym, error) {
	var gym []models.Gym
	err := db.Conn.Find(&gym).Error
	return gym, err
}

func (db *dbGym) CreateNewGym(gym models.Gym) error {
	return db.Conn.Create(&gym).Error
}

func (db *dbGym) UpdateGym(gym models.Gym) error {
	return db.Conn.Save(&gym).Error
}

func (db *dbGym) DeleteGym(idGym uuid.UUID) error {
	return db.Conn.Delete(&models.Gym{}, idGym).Error
}

func (db *dbGym) GetGymById(idGym uuid.UUID) (models.Gym, error) {
	var gym models.Gym
	err := db.Conn.Where("id_gym = ?", idGym).First(&gym).Error
	return gym, err
}

type GymRepository interface {
	GetGym() ([]models.Gym, error)
	CreateNewGym(gym models.Gym) error
	UpdateGym(gym models.Gym) error
	DeleteGym(idGym uuid.UUID) error
	GetGymById(idGym uuid.UUID) (models.Gym, error)
}

func NewDBGymRepository(conn *gorm.DB) *dbGym {
	return &dbGym{Conn: conn}
}
