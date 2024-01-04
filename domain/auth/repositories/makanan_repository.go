package repositories

import (
	"kalorize-api/domain/auth/models"

	"gorm.io/gorm"
)

type dbMakanan struct {
	Conn *gorm.DB
}

func (db *dbMakanan) GetAllMakanan() ([]models.Makanan, error) {
	var makanans []models.Makanan
	err := db.Conn.Find(&makanans).Error
	return makanans, err
}

func (db *dbMakanan) GetMakananById(id string) (models.Makanan, error) {
	var makanan models.Makanan
	err := db.Conn.Where("id_makanan = ?", id).First(&makanan).Error
	return makanan, err
}

func (db *dbMakanan) CreateMakanan(makanan models.Makanan) error {
	return db.Conn.Create(&makanan).Error
}

type MakananRepository interface {
	GetAllMakanan() ([]models.Makanan, error)
	GetMakananById(id string) (models.Makanan, error)
	CreateMakanan(makanan models.Makanan) error
}

func NewDBMakananRepository(conn *gorm.DB) *dbMakanan {
	return &dbMakanan{Conn: conn}
}
