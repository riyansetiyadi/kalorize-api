package models

import (
	"time"

	"github.com/google/uuid"
)

type KodeGym struct {
	IdKodeGym   uuid.UUID `json:"id_kode" gorm:"column:id_kode;primary_key;type:int(12);"`
	KodeGym     string    `json:"kode_gym" gorm:"column:kode_gym;type:varchar(255);"` //misal "bojong56"
	IdGym       uuid.UUID `json:"id_gym" gorm:"column:id_gym;type:char(36);"`
	ExpiredTime time.Time `json:"expired_date" gorm:"column:expired_date;type:timestamp;"`
}
