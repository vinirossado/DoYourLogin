package repository

import (
	"abrigos/source/configuration"
	"abrigos/source/domain/entities"
	"abrigos/source/domain/enumerations"
)

func AutoMigrate() {
	db.AutoMigrate(entities.RetriveAll()...)
	// InsertBaseUsers()
}

func DropAll() {
	db.Migrator().DropTable(entities.RetriveAll()...)
}

func InsertBaseUsers() {
	if configuration.APPLICATION_ENVIRONMENT.ValueAsString() == "dev" {
		CreateUser(&entities.User{
			Name:     "admin",
			Username: "admin",
			Password: "admin",
			Role:     enumerations.ADMIN,
		})
	}
}
