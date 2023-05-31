package entities

import (
	"doYourLogin/source/domain/enumerations"

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

	CompanyID uint `gorm:"column:company_id"`

	Company Company `gorm:"foreignKey:CompanyID"`
}

func NewUser(name, username, email, password, address, phone, about, image string, role enumerations.Roles, companyID uint) *User {
	passwordHashed, err := BeforeCreate(password)

	if err != nil {
		panic(err)
	}

	user := User{
		Name:      name,
		Username:  username,
		Email:     email,
		Address:   address,
		Phone:     phone,
		About:     about,
		Image:     image,
		Password:  string(passwordHashed),
		Role:      role,
		CompanyID: companyID,
	}

	return &user
}

func BeforeCreate(password string) (hash []byte, error error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return hash, nil
}
