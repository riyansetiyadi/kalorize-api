package models

import (
	"time"

	"github.com/google/uuid"
)

type History struct {
	IdHistory     uuid.UUID `json:"id_history" gorm:"column:id_history;primary_key;auto_increment;"`
	IdUser        uuid.UUID `json:"id_user" gorm:"column:id_user;type:char(16);"`
	IdBreakfast   string    `json:"id_breakfast" gorm:"column:id_breakfast;type:char(36);"`
	IdLunch       string    `json:"id_lunch" gorm:"column:id_lunch;type:char(36);"`
	IdDinner      string    `json:"id_dinner" gorm:"column:id_dinner;type:char(36);"`
	TotalProtein  int       `json:"total_protein" gorm:"column:total_protein;type:int(11);"`
	TotalKalori   int       `json:"total_kalori" gorm:"column:total_kalori;type:int(11);"`
	TanggalDibuat time.Time `json:"tanggal_dibuat" gorm:"column:tanggal_dibuat;type:datetime;"`
}
