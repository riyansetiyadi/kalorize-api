package models

import (
	"time"

	"github.com/google/uuid"
)

type Makanan struct {
	ID          uuid.UUID `json:"id_makanan" gorm:"column:id_makanan;primary_key;type:char(36);"`
	Nama        string    `json:"nama" gorm:"column:nama;type:varchar(255);"`
	Kalori      int       `json:"kalori" gorm:"column:kalori;type:int;"`
	Protein     int       `json:"protein" gorm:"column:protein;type:int;"`
	Bahan       string    `json:"bahan" gorm:"column:bahan;type:varchar(255);"`
	CookingStep string    `json:"cooking_step" gorm:"column:cooking_step;type:varchar(255);"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
