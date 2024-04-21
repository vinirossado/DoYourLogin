package repositories

import (
	"doYourLogin/source/domain/entities"
)

func FindUsers(tx ...*TransactionalOperation) ([]entities.User, error) {
	var users []entities.User
	return users, WithTransactions(tx).Find(&users).Error
}

func FindUserByUsername(username string, tx ...*TransactionalOperation) (*entities.User, error) {
	user := entities.User{}
	return &user, WithTransactions(tx).Where("username = ?", username).First(&user).Error
}

func FindUserByEmail(email string, tx ...*TransactionalOperation) (*entities.User, error) {
	user := entities.User{}
	return &user, WithTransactions(tx).Where("email = ?", email).First(&user).Error
}

func ExistsUserByUsernameOrEmail(username, email string, tx ...*TransactionalOperation) bool {
	dbResult := WithTransactions(tx).Where("username = ? OR email = ?", username, email).Find(&entities.User{})
	return dbResult.RowsAffected > 0
}

func FindUserById(id int, tx ...*TransactionalOperation) (*entities.User, error) {
	user := entities.User{}
	return &user, WithTransactions(tx).Where("id = ?", id).First(&user).Error
}

func CreateUser(user *entities.User, tx ...*TransactionalOperation) error {
	return WithTransactions(tx).Create(user).Error
}

func UpdateUser(user *entities.User, tx ...*TransactionalOperation) error {
	return WithTransactions(tx).Updates(&user).Error
}

func DeleteUser(id int, tx ...*TransactionalOperation) error {
	//dbResult := WithTransactions(tx).Delete(&entities.User{Model: gorm.Model{ID: uint(id)}})
	//
	//if dbResult.RowsAffected < 1 {
	//	return gorm.ErrRecordNotFound
	//}
	//
	//return dbResult.Error
	return nil
}
