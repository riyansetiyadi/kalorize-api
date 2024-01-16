package models

import "github.com/google/uuid"

type UsedCode struct {
	IdGym   uuid.UUID `json:"id_gym" gorm:"column:id_gym;primary_key;type:char(36);"`
	KodeGym string    `json:"kode_gym" gorm:"column:id_kode;primary_key;type:char(36);"`
	IdUser  uuid.UUID `json:"id_user" gorm:"column:id_user;type:char(36);"`
}

func (UsedCode) TableName() string {
	return "used_codes"
}
