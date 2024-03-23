package models

import "github.com/google/uuid"

type Franchise struct {
	IdFranchise        uuid.UUID `json:"id_franchise" gorm:"column:id_franchise;primary_key; type:char(36);"`
	NamaFranchise      string    `json:"nama_franchise" gorm:"column:nama_franchise;type:varchar(255);"`
	LongitudeFranchise float64   `json:"longitude_franchise" gorm:"column:longitude_franchise;type:double;"`
	LatitudeFranchise  float64   `json:"latitude_franchise" gorm:"column:latitude_franchise;type:double;"`
	NoTeleponFranchise string    `json:"telepon_franchise" gorm:"column:telepon;type:varchar(16);"`
	FotoFranchise      string    `json:"foto_franchise" gorm:"column:foto;type:varchar(255);"`
	EmailFranchise     string    `json:"email_franchise" gorm:"column:email;type:varchar(255);"`
	PasswordFranchise  string    `json:"password_franchise" gorm:"column:password;type:varchar(255);"`
	LokasiFranchise    string    `json:"lokasi_franchise" gorm:"column:lokasi;type:varchar(255);"`
}

func (Franchise) TableName() string {
	return "franchises"
}
