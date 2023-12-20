package models

import "time"

type MealSet struct {
	IDUser         int       `json:"id_user" gorm:"column:id_user"`
	IDMakanan      int       `json:"id_makanan" gorm:"column:id_makanan"`
	JumlahKalori   int       `json:"jumlah_kalori" gorm:"column:jumlah_kalori"`
	JumlahProtein  int       `json:"jumlah_protein" gorm:"column:jumlah_protein"`
	TanggalMealSet time.Time `json:"tanggal_meal_set" gorm:"column:tanggal_meal_set"`
}
