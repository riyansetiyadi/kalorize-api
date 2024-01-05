package models

import "github.com/google/uuid"

type UsedCode struct {
	IdKode uuid.UUID `json:"id_kode" gorm:"column:id_kode;primary_key;type:char(36);"`
	IdUser uuid.UUID `json:"id_user" gorm:"column:id_user;type:char(36);"`
}
