package entities

import (
	"gorm.io/gorm"
	"time"
)

type Company struct {
	gorm.Model
	Name            string    `gorm:"column:name;unique"`
	APIToken        string    `gorm:"column:apiToken;unique"`
	Active          bool      `gorm:"column:active"`
	ExpirationToken time.Time `gorm:"column:expirationToken"`

	Users []User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
