package models

import (
	"time"
)

type Makanan struct {
	ID          int       `json:"id" gorm:"column:id"`
	Nama        string    `json:"nama" gorm:"column:nama"`
	Kalori      int       `json:"kalori" gorm:"column:kalori"`
	Protein     int       `json:"protein" gorm:"column:protein"`
	Bahan       string    `json:"bahan" gorm:"column:bahan"`
	CookingStep string    `json:"cooking_step" gorm:"column:cooking_step"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}
