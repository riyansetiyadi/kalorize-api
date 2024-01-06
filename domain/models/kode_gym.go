package models

import (
	"time"
)

type KodeGym struct {
	IdKodeGym   int       `json:"id_kode" gorm:"column:id_kode;primary_key;type:int(12);"`
	KodeGym     string    `json:"kode_gym" gorm:"column:kode_gym;type:varchar(255);"`
	IdGym       int       `json:"id_gym" gorm:"column:id_gym;type:int(12);"`
	ExpiredTime time.Time `json:"expired_date" gorm:"column:expired_date;type:timestamp;"`
}
