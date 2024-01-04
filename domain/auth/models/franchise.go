package models

type Franchise struct {
	IdFranchise   int    `json:"id_franchise" gorm:"column:id_franchise;primary_key;auto_increment;"`
	NamaFranchise string `json:"nama_franchise" gorm:"column:nama_franchise;type:varchar(255);"`
	Alamat        string `json:"alamat" gorm:"column:alamat;type:varchar(255);"`
}
