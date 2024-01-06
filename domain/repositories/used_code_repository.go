package repositories

import (
	"kalorize-api/domain/models"

	"gorm.io/gorm"
)

type UsedCode struct {
	Conn *gorm.DB
}

func (db *UsedCode) GetUsedCode() ([]models.UsedCode, error) {
	var useCode []models.UsedCode
	err := db.Conn.Find(&useCode).Error
	return useCode, err
}

func (db *UsedCode) CreateNewUsedCode(useCode models.UsedCode) error {
	return db.Conn.Create(&useCode).Error
}

func (db *UsedCode) UpdateUsedCode(useCode models.UsedCode) error {
	return db.Conn.Save(&useCode).Error
}

func (db *UsedCode) DeleteUsedCode(idUsedCode int) error {
	return db.Conn.Delete(&models.UsedCode{}, idUsedCode).Error
}

func (db *UsedCode) GetUsedCodeByIdCode(idUsedCode int) (models.UsedCode, error) {
	var usedCode models.UsedCode
	err := db.Conn.Where(" id_kode = ?", idUsedCode).First(&usedCode).Error
	return usedCode, err
}

func (db *UsedCode) GetusedCodeByIdUser(idUser int) (models.UsedCode, error) {
	var usedCode models.UsedCode
	err := db.Conn.Where(" id_user = ?", idUser).First(&usedCode).Error
	return usedCode, err
}

type UsedCodeRepository interface {
	GetUsedCode() ([]models.UsedCode, error)
	CreateNewUsedCode(useCode models.UsedCode) error
	UpdateUsedCode(models.UsedCode) error
	DeleteUsedCode(idUseCode int) error
	GetUsedCodeByIdCode(idUseCode int) (models.UsedCode, error)
	GetusedCodeByIdUser(idUser int) (models.UsedCode, error)
}

func NewDBUsedCodeRepository(conn *gorm.DB) *UsedCode {
	return &UsedCode{Conn: conn}
}
