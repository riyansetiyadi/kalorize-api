package models

type Franchise struct {
	IdFranchise      string `json:"id_franchise" gorm:"column:id_franchise;primary_key;auto_increment;"`
	NamaFranchise    string `json:"nama_franchise" gorm:"column:nama_franchise;type:varchar(255);"`
	AlamatFranchise  string `json:"alamat_franchise" gorm:"column:alamat;type:varchar(255);"`
	TeleponFranchise string `json:"telepon_franchise" gorm:"column:telepon;type:varchar(16);"`
	FotoFranchise    string `json:"foto_franchise" gorm:"column:foto;type:varchar(255);"`
	LokasiFranchise  string `json:"lokasi_franchise" gorm:"column:lokasi;type:varchar(255);"`
}

func (Franchise) TableName() string {
	return "franchises"
}
