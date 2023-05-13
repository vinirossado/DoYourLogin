package repositories

import (
	"doYourLogin/source/configuration"
	"doYourLogin/source/domain/entities"
	"doYourLogin/source/domain/enumerations"
)

func AutoMigrate() {
	db.AutoMigrate(entities.RetriveAll()...)

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
			Email:    "vinirossado@gmail.com",
			Role:     enumerations.GOD,
		})
	}
}
