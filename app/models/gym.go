package models

import (
	"github.com/google/uuid"
)

type Gym struct {
	IdGym      uuid.UUID `json:"id" gorm:"column:id;primary_key;type:char(36);"`
	NamaGym    string    `json:"nama" gorm:"column:nama;type:varchar(255);"`
	AlamatGym  string    `json:"alamat" gorm:"column:alamat;type:varchar(255);"`
	Latitude   float64   `json:"latitude" gorm:"column:latitude;type:double;"`
	Longitude  float64   `json:"longitude" gorm:"column:longitude;type:double;"`
	LinkGoogle string    `json:"link_google" gorm:"column:link_google;type:varchar(255);"`
}

func (Gym) TableName() string {
	return "gyms"
}
