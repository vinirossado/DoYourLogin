package entities

import (
	"doYourLogin/source/domain/enumerations"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID uint `gorm:"primarykey"`

	Name      string             `gorm:"column:name"`
	Username  string             `gorm:"column:username;unique"`
	Email     string             `gorm:"column:email;unique"`
	Address   string             `gorm:"column:address"`
	Phone     string             `gorm:"column:phone"`
	About     string             `gorm:"column:about"`
	Image     string             `gorm:"column:image"`
	Password  string             `gorm:"column:password"`
	Active    bool               `gorm:"column:active"`
	Role      enumerations.Roles `gorm:"column:role"`
	CreatedAt time.Time          `gorm:"created_at"`
	UpdatedAt time.Time          `gorm:"updated_at"`
	DeletedAt time.Time          `gorm:"index"`
}

func NewUser(name, username, email, password, address, phone, about, image string, role enumerations.Roles) *User {
	passwordHashed, err := BeforeCreate(password)

	if err != nil {
		panic(err)
	}

	user := User{
		Name:     name,
		Username: username,
		Email:    email,
		Address:  address,
		Phone:    phone,
		About:    about,
		Image:    image,
		Password: string(passwordHashed),
		Role:     role,
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
