package repositories

import (
	"doYourLogin/source/configuration"
	"doYourLogin/source/domain/entities"
	"doYourLogin/source/domain/enumerations"
)

func AutoMigrate() {
	err := db.AutoMigrate(entities.RetriveAll()...)
	if err != nil {
		return
	}

}

func DropAll() {
	db.Migrator().DropTable(entities.RetriveAll()...)
}

func InsertBaseUsers() {
	if configuration.APPLICATION_ENVIRONMENT.ValueAsString() == "dev" {
		err := CreateUser(&entities.User{
			Name:     "admin",
			Username: "admin",
			Password: "admin",
			Email:    "vinirossado@gmail.com",
			Role:     enumerations.GOD,
		})
		if err != nil {
			return
		}
	}
}
