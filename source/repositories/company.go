package repositories

import "doYourLogin/source/domain/entities"

func CreateCompany(company *entities.Company, tx ...*TransactionalOperation) (uint, error) {
	err := WithTransaction(tx).Create(company).Error
	return company.ID, err
}
