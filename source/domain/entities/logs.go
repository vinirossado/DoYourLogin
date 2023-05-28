package entities

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Logs struct {
	gorm.Model
	UserID    uint            `gorm:"column:user_id"`
	CompanyID uint            `gorm:"column:company_id"`
	Error     string          `gorm:"column:error"`
	Message   string          `gorm:"column:message"`
	Route     string          `gorm:"column:route"`
	Source    string          `gorm:"column:source"`
	Level     logger.LogLevel `gorm:"column:level"`
	Timestamp time.Time       `gorm:"column:timestamp"`

	Company Company `gorm:"foreignKey:CompanyID"`
	User    User    `gorm:"foreignKey:UserID"`
}
