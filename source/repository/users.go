package repository

import (
	"abrigos/source/domain/entities"

	"gorm.io/gorm"
)

func FindUsers(tx ...*TransactionalOperation) ([]entities.User, error) {
	users := []entities.User{}
	return users, WithTransaction(tx).Find(&users).Error
}

func FindUserByUsername(username string, tx ...*TransactionalOperation) (*entities.User, error) {
	user := &entities.User{}
	return user, WithTransaction(tx).Where("username = ?", username).First(user).Error
}

func FindUserByEmail(email string, tx ...*TransactionalOperation) (*entities.User, error) {
	user := &entities.User{}
	return user, WithTransaction(tx).Where("email = ?", email).First(user).Error
}

func ExistsUserByUsername(username string, tx ...*TransactionalOperation) bool {
	dbResult := WithTransaction(tx).Where("username = ?", username).Find(&entities.User{})
	return dbResult.RowsAffected > 0
}

func FindUserById(id int, tx ...*TransactionalOperation) (*entities.User, error) {
	user := &entities.User{}
	return user, WithTransaction(tx).Where("id = ?", id).First(user).Error
}

func CreateUser(user *entities.User, tx ...*TransactionalOperation) error {
	return WithTransaction(tx).Create(user).Error
}

func UpdateUser(user *entities.User, tx ...*TransactionalOperation) error {
	// return WithTransaction(tx).Model(&entities.User{}).Where("id = ?", user.ID).Updates(user).Error
	return WithTransaction(tx).Updates(&user).Error

}

func DeleteUser(id int, tx ...*TransactionalOperation) error {
	dbResult := WithTransaction(tx).Delete(&entities.User{Model: gorm.Model{ID: uint(id)}})

	if dbResult.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return dbResult.Error
}
