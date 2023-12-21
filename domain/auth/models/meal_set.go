package models

import (
	"time"

	"github.com/google/uuid"
)

type MealSet struct {
	IdUser         uuid.UUID `json:"id_user" gorm:"column:id_user;type:char(36);"`
	IdMakanan      uuid.UUID `json:"id_makanan" gorm:"column:id_makanan;type:char(36);"`
	JumlahKalori   int       `json:"jumlah_kalori" gorm:"column:jumlah_kalori;type:int;"`
	JumlahProtein  int       `json:"jumlah_protein" gorm:"column:jumlah_protein;type:int;"`
	TanggalMealSet time.Time `json:"tanggal_meal_set" gorm:"column:tanggal_meal_set;type:timestamp;"`
}
