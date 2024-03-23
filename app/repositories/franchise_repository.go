package repositories

import (
	"kalorize-api/app/models"

	"gorm.io/gorm"
)

type dbFranchise struct {
	Conn *gorm.DB
}

func (db *dbFranchise) GetAllFranchise() ([]models.Franchise, error) {
	var franchises []models.Franchise
	err := db.Conn.Find(&franchises).Error
	return franchises, err
}

func (db *dbFranchise) GetFranchiseById(id string) (models.Franchise, error) {
	var franchise models.Franchise
	err := db.Conn.Where("id_franchise = ?", id).First(&franchise).Error
	return franchise, err
}

func (db *dbFranchise) CreateFranchise(franchise models.Franchise) error {
	return db.Conn.Create(&franchise).Error
}

func (db *dbFranchise) DeleteFranchise(franchise models.Franchise) error {
	return db.Conn.Delete(&franchise).Error
}

func (db *dbFranchise) AddFranchiseMakanan(franchiseMakanan models.FranchiseMakanan) error {
	return db.Conn.Create(&franchiseMakanan).Error
}

func (db *dbFranchise) UpdateFranchise(franchise models.Franchise) error {
	return db.Conn.Save(&franchise).Error
}

type FranchiseRepository interface {
	UpdateFranchise(franchise models.Franchise) error
	AddFranchiseMakanan(franchiseMakanan models.FranchiseMakanan) error
	GetAllFranchise() ([]models.Franchise, error)
	GetFranchiseById(id string) (models.Franchise, error)
	CreateFranchise(franchise models.Franchise) error
}

func NewDBFranchiseRepository(conn *gorm.DB) *dbFranchise {
	return &dbFranchise{Conn: conn}
}
