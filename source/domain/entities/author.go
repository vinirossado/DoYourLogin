package entities

import "time"

type Author struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `gorm:"column:name;unique"`
	Email     string    `gorm:"column:email;unique"`
	Active    bool      `gorm:"column:active"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	DeletedAt time.Time `gorm:"index"`
}
