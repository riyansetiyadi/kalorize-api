package models

import (
	"github.com/google/uuid"
)

type Gym struct {
	IdGym   uuid.UUID `json:"id_gym" gorm:"column:id_gym;primary_key;type:char(36);"`
	NamaGym string    `json:"nama_gym" gorm:"column:nama_gym;type:varchar(255);"`
	Alamat  string    `json:"alamat" gorm:"column:alamat;type:varchar(255);"`
}
