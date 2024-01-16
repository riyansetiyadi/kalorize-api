package models

import "github.com/google/uuid"

type FranchiseMakanan struct {
	IdFranchiseMakanan uuid.UUID `json:"id_franchise_makanan" gorm:"column:id_franchise_makanan;primary_key;auto_increment; type:char(36);"`
	IdFranchise        uuid.UUID `json:"id_franchise" gorm:"column:id_franchise;type:char(36);"`
	IdMakanan          string    `json:"id_makanan" gorm:"column:id_makanan;type:char(36);"`
}

func (FranchiseMakanan) TableName() string {
	return "franchise_makanans"
}
