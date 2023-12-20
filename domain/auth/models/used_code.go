package models

type UsedCode struct {
	IdKode string `json:"id_kode" gorm:"column:id_kode"`
	IdUser string `json:"id_user" gorm:"column:id_user"`
}
