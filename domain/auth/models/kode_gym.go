package models

import "time"

type KodeGym struct {
	IdKodeGym   int       `json:"id_kode_gym" gorm:"column:id_kode_gym"`
	KodeGym     string    `json:"kode_gym" gorm:"column:kode_gym"`
	ExpiredTime time.Time `json:"expired_date" gorm:"column:expired_date"`
}
