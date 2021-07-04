package database

import (
	"LazarusPoC/configs"
	"LazarusPoC/helpers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnDB() {
	db, err := gorm.Open(postgres.Open(configs.PG_URI), &gorm.Config{})
	helpers.HandleErr(err)

	// db.AutoMigrate(&models.User{}, &models.Organization{}, &models.Membership{})

	DB = db
}
