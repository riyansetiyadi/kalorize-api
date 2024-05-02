package models

import (
	"github.com/google/uuid"
)

type Gym struct {
	IdGym      uuid.UUID `json:"id_gym" gorm:"column:id_gym;primary_key;type:char(36);"`
	NamaGym    string    `json:"nama_gym" gorm:"column:nama_gym;type:varchar(255);"`
	AlamatGym  string    `json:"alamat_gym" gorm:"column:alamat_gym;type:varchar(255);"`
	Latitude   float64   `json:"latitude" gorm:"column:latitude;type:double;"`
	Longitude  float64   `json:"longitude" gorm:"column:longitude;type:double;"`
	LinkGoogle string    `json:"link_google" gorm:"column:link_google;type:varchar(255);"`
}

func (Gym) TableName() string {
	return "gyms"
}
