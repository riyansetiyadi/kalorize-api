package models

import "github.com/google/uuid"

type Token struct {
	IdToken      uuid.UUID `json:"id_item" gorm:"column:id_token;type:char(36);primary_key"`
	UserId       uuid.UUID `json:"user_id" gorm:"column:user_id;type:char(36)"`
	AccessToken  string    `json:"access_token" gorm:" column:access_token;type:varchar(255);"`
	RefreshToken string    `json:"refresh_token" gorm:" column:refresh_token;type:varchar(255);"`
}

func (t *Token) TableName() string {
	return "tokens"
}
