package entities

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name     string `gorm:"column:name;unique"`
	APIToken string `gorm:"column:apiToken;unique"`

	Users []User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
