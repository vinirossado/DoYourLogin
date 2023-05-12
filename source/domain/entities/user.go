package entities

import (
	"abrigos/source/domain/enumerations"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string             `gorm:"column:name"`
	Username string             `gorm:"column:username;unique"`
	Email    string             `gorm:"column:email;unique"`
	Address  string             `gorm:"column:address"`
	Phone    string             `gorm:"column:phone"`
	About    string             `gorm:"column:about"`
	Image    string             `gorm:"column:image"`
	Password string             `gorm:"column:password"`
	Role     enumerations.Roles `gorm:"column:role"`
}

func (u *User) BeforeCreate(scope *gorm.DB) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	scope.Table("users", "password", hash)
	return nil
}
