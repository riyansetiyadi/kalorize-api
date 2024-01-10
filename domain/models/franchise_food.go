package models

type FranchiseMakanan struct {
	IdFranchiseMakanan string `json:"id_franchise_makanan" gorm:"column:id_franchise_makanan;primary_key;auto_increment;"`
	IdFranchise        string `json:"id_franchise" gorm:"column:id_franchise;type:string(11);"`
	IdMakanan          string `json:"id_makanan" gorm:"column:id_makanan;type:int(11);"`
}

func (FranchiseMakanan) TableName() string {
	return "franchise_makanans"
}
