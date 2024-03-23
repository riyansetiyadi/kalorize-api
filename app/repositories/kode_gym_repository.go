package repositories

import (
	"kalorize-api/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbKodeGym struct {
	Conn *gorm.DB
}

func (db *dbKodeGym) GetKodeGymByKode(kode string) (models.KodeGym, error) {
	var kodeGym models.KodeGym
	err := db.Conn.Where("kode_gym = ?", kode).First(&kodeGym).Error
	if err != nil {
		emptyUUID := uuid.UUID{}
		kodeGym.IdKodeGym = emptyUUID
	}
	return kodeGym, err
}

func (db *dbKodeGym) GetIDFromKode(kode string) (uuid.UUID, error) {
	var kodeGym models.KodeGym
	err := db.Conn.Where("kode_gym = ?", kode).First(&kodeGym).Error
	return kodeGym.IdKodeGym, err
}

func (db *dbKodeGym) CreateNewKodeGym(kodeGym models.KodeGym) error {
	return db.Conn.Create(&kodeGym).Error
}

func (db *dbKodeGym) UpdateKodeGym(kodeGym models.KodeGym) error {
	return db.Conn.Save(&kodeGym).Error
}

func (db *dbKodeGym) DeleteKodeGym(idKodeGym uuid.UUID) error {
	return db.Conn.Delete(&models.KodeGym{}, idKodeGym).Error
}

func (db *dbKodeGym) GetKodeGymById(idKodeGym uuid.UUID) (models.KodeGym, error) {
	var kodeGym models.KodeGym
	err := db.Conn.Where("id_kode_gym = ?", idKodeGym).First(&kodeGym).Error
	return kodeGym, err
}

type KodeGymRepository interface {
	GetKodeGymByKode(kode string) (models.KodeGym, error)
	CreateNewKodeGym(kodeGym models.KodeGym) error
	UpdateKodeGym(kodeGym models.KodeGym) error
	DeleteKodeGym(idKodeGym uuid.UUID) error
	GetKodeGymById(idKodeGym uuid.UUID) (models.KodeGym, error)
	GetIDFromKode(kode string) (uuid.UUID, error)
}

func NewDBKodeGymRepository(conn *gorm.DB) *dbKodeGym {
	return &dbKodeGym{Conn: conn}
}
