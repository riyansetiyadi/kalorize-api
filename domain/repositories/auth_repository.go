package repositories

import (
	"kalorize-api/domain/models"

	"gorm.io/gorm"
)

type dbAuth struct {
	Conn *gorm.DB
}

func (dbAuth *dbAuth) GetToken() string {
	return "token"
}

func (db *dbAuth) CreateNewUser(user models.User) error {
	return db.Conn.Create(&user).Error
}

func (db *dbAuth) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := db.Conn.Where("username = ?", username).First(&user).Error
	return user, err
}

func (db *dbAuth) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := db.Conn.Where("email = ?", email).First(&user).Error
	return user, err
}

type AuthRepository interface {
	GetToken() string
	CreateNewUser(user models.User) error
	GetUserByUsername(username string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
}

func NewDBAuthRepository(conn *gorm.DB) *dbAuth {
	return &dbAuth{Conn: conn}
}
