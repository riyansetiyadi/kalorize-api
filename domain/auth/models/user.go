package models

type User struct {
	IdUser       int    `json:"id" gorm:"column:id"`
	Fullname     string `json:"fullname" gorm:"column:full_name"`
	Email        string `json:"email" gorm:"column:email"`
	Password     string `json:"password" gorm:"column:password"`
	JenisKelamin string `json:"jenis_kelamin" gorm:"column:jenis_kelamin"`
	Umur         int    `json:"umur" gorm:"column:umur"`
	BeratBadan   int    `json:"berat_badan" gorm:"column:berat_badan"`
	TinggiBadan  int    `json:"tinggi_badan" gorm:"column:tinggi_badan"`
	FrekuensiGym int    `json:"frekuensi_gym" gorm:"column:frekuensi_gym"`
	TargetKalori int    `json:"target_kalori" gorm:"column:target_kalori"`
}

func (u *User) TableName() string {
	return "users"
}
