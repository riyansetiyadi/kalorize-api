package models

import "github.com/google/uuid"

type TargetKalori struct {
}

type User struct {
	IdUser       uuid.UUID `json:"id_user" gorm:"column:id_user;primary_key;type:char(36);"`
	Fullname     string    `json:"fullname" gorm:"column:full_name;type:varchar(255);"`
	Email        string    `json:"email" gorm:"column:email;type:varchar(255);"`
	Password     string    `json:"password" gorm:"column:password;type:varchar(255);"`
	Role         string    `json:"role" gorm:"column:role;type:varchar(20);"`
	JenisKelamin int       `json:"jenis_kelamin" gorm:"column:jenis_kelamin;type:int(2);"`
	Umur         int       `json:"umur" gorm:"column:umur;type:int;"`
	BeratBadan   int       `json:"berat_badan" gorm:"column:berat_badan;type:int;"`
	TinggiBadan  int       `json:"tinggi_badan" gorm:"column:tinggi_badan;type:int;"`
	FrekuensiGym int       `json:"frekuensi_gym" gorm:"column:frekuensi_gym;type:int;"`
	TargetKalori int       `json:"target_kalori" gorm:"column:target_kalori;type:int;"`
	ReferalCode  string    `json:"referal_code" gorm:"column:referal_code;type:varchar(255);"`
}

func (u *User) TableName() string {
	return "users"
}
