package models

type Gym struct {
	IdGym   int    `json:"id_gym" gorm:"column:id_gym"`
	NamaGym string `json:"nama_gym" gorm:"column:nama_gym"`
	Alamat  string `json:"alamat" gorm:"column:alamat"`
}
