package models

type Token struct {
	IdToken int64  `json:"id_item" gorm:"column:id_item;primaryKey;autoIncrement"`
	UserId  int64  `json:"user_id" gorm:"type:int;"`
	Token   string `json:"token" gorm:"type:string;"`
}

func (t *Token) TableName() string {
	return "tokens"
}
