package entities

import (
	"gorm.io/gorm"
	"time"
)

type Logs struct {
	gorm.Model
	UserID    uint      `gorm:"column:user_id"`
	CompanyID uint      `gorm:"column:company_id"`
	Error     uint      `gorm:"column:error;type:string"`
	Message   string    `gorm:"column:message"`
	Route     string    `gorm:"column:route"`
	Source    string    `gorm:"column:source"`
	Level     uint      `gorm:"column:level"`
	Timestamp time.Time `gorm:"column:timestamp"`

	Company Company `gorm:"foreignKey:CompanyID"`
	User    User    `gorm:"foreignKey:UserID"`
}
