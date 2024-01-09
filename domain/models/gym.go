package models

import (
	"github.com/google/uuid"
)

type Gym struct {
	IdGym        uuid.UUID `json:"id_gym" gorm:"column:id_gym;primary_key;type:char(36);"`
	NamaGym      string    `json:"nama_gym" gorm:"column:nama_gym;type:varchar(255);"`
	AlamatGym    string    `json:"alamat_gym" gorm:"column:alamat;type:varchar(255);"`
	EmailGym     string    `json:"email_gym" gorm:"column:email;type:varchar(255);"`
	NoTeleponGym string    `json:"no_telepon_gym" gorm:"column:no_telepon;type:varchar(255);"`
	PasswordGym  string    `json:"password_gym" gorm:"column:password;type:varchar(255);"`
}

func (Gym) TableName() string {
	return "gyms"
}
