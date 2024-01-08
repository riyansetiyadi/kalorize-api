package models

import "github.com/google/uuid"

type UsedCode struct {
	IdUsedCode string    `json:"id_used_code" gorm:"column:id_used_code;primary_key;type:char(36);"`
	KodeGym    string    `json:"kode_gym" gorm:"column:id_kode;primary_key;type:char(36);"`
	IdUser     uuid.UUID `json:"id_user" gorm:"column:id_user;type:char(36);"`
}

func (UsedCode) TableName() string {
	return "used_codes"
}
