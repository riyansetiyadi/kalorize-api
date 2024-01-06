package models

type UsedCode struct {
	IdUsedCode int `json:"id_used_code" gorm:"column:id_used_code;primary_key;type:int(10);"`
	IdKode     int `json:"id_kode" gorm:"column:id_kode;primary_key;type:int(10);"`
	IdUser     int `json:"id_user" gorm:"column:id_user;type:int(10);"`
}

func (UsedCode) TableName() string {
	return "used_codes"
}
