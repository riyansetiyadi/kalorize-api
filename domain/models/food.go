package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Makanan struct {
	IdMakanan   string      `json:"id" gorm:"column:id;primary_key;type:char(36);"`
	Nama        string      `json:"nama" gorm:"column:nama;type:varchar(255);"`
	Jenis       string      `json:"jenis" gorm:"column:jenis;type:varchar(255);"`
	Foto        string      `json:"foto" gorm:"column:foto;type:varchar(255);"`
	Kalori      int         `json:"kalori" gorm:"column:kalori;type:int;"`
	Protein     int         `json:"protein" gorm:"column:protein;type:int;"`
	Bahan       string      `json:"bahan" gorm:"column:bahan;type:text;"`
	CookingStep string      `json:"cooking_step" gorm:"column:cooking_step;type:text;"`
	CreatedAt   TimeWrapper `json:"created_at" gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   TimeWrapper `json:"updated_at" gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (m *Makanan) TableName() string {
	return "makanans"
}

type TimeWrapper struct {
	time.Time
}

func (tw *TimeWrapper) Scan(value interface{}) error {
	if value == nil {
		tw.Time = time.Time{}
		return nil
	}

	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("failed to scan time")
	}

	tw.Time = t
	return nil
}

func (tw TimeWrapper) Value() (driver.Value, error) {
	return tw.Time, nil
}
