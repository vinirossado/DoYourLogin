package utils

import (
	"doYourLogin/source/domain/entities"
	"doYourLogin/source/domain/enumerations"
	"doYourLogin/source/domain/requests"
	"encoding/json"
	"fmt"
	"testing"
)

type Source struct {
	Field1 string
	Field2 int
}

type Destination struct {
	FieldA string
	FieldB int
}

func TestMap(test *testing.T) {

	user := entities.User{
		Name:     "ABC",
		Username: "user1",
		Email:    "dycjh@example.com",
		Address:  "address1",
		About:    "about1",
		Password: "password1",
		Image:    "image1",
		Phone:    "phone1",
		Role:     enumerations.GOD,
	}

	dto := requests.UserRequest{}

	Map(user, &dto)

	jsonData, err := json.Marshal(dto)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonData))

}
