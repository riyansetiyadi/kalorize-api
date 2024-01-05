package config

import (
	"kalorize-api/domain/models"

	"gorm.io/gorm"
)

func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Token{})
	db.AutoMigrate(&models.UsedCode{})
	db.AutoMigrate(&models.Gym{})
	db.AutoMigrate(&models.Makanan{})
	db.AutoMigrate(&models.KodeGym{})
	db.AutoMigrate(&models.MealSet{})
}
