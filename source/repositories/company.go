package repositories

import "doYourLogin/source/domain/entities"

func CreateCompany(company *entities.Company, tx ...*TransactionalOperation) (uint, error) {
	err := WithTransactions(tx).Create(company).Error
	return company.ID, err
}

func FindCompanyById(id uint, tx ...*TransactionalOperation) (*entities.Company, error) {
	company := entities.Company{}
	return &company, WithTransactions(tx).Where("id =?", id).First(&company).Error
}

func FindCompanyByApiToken(apiToken string, tx ...*TransactionalOperation) (*entities.Company, error) {
	company := entities.Company{}
	return &company, WithTransactions(tx).Where("apiToken =?", apiToken).First(&company).Error
}

func ActivateAccount(company *entities.Company, tx ...*TransactionalOperation) error {
	return WithTransactions(tx).Updates(&company).Error
}

func FindCompanies(tx ...*TransactionalOperation) ([]entities.Company, error) {
	var companies []entities.Company
	return companies, WithTransactions(tx).Find(&companies).Error
}
