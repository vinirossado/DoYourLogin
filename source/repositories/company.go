package repositories

import "doYourLogin/source/domain/entities"

func CreateCompany(company *entities.Company, tx ...*TransactionalOperation) (uint, error) {
	err := WithTransaction(tx).Create(company).Error
	return company.ID, err
}

func FindCompanyById(id int, tx ...*TransactionalOperation) (*entities.Company, error) {
	company := &entities.Company{}
	return company, WithTransaction(tx).Where("id =?", id).First(company).Error
}

func FindCompanies(tx ...*TransactionalOperation) ([]entities.Company, error) {
	companies := []entities.Company{}
	return companies, WithTransaction(tx).Find(&companies).Error
}
