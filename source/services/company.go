package services

import (
	"crypto/rand"
	"doYourLogin/source/domain/entities"
	"doYourLogin/source/domain/enumerations"
	"doYourLogin/source/domain/exceptions"
	"doYourLogin/source/domain/requests"
	"doYourLogin/source/domain/responses"
	"doYourLogin/source/repositories"
	"fmt"
)

func CreateCompany(request *requests.CompanyRequest) responses.CompanyResponse {
	var companyID uint
	var apiToken string

	repositories.UsingTransactional(func(tx *repositories.TransactionalOperation) error {
		tx.BeginTransaction()

		if repositories.ExistsUserByUsername(request.Name) {
			return fmt.Errorf("Company %s already exists", request.Name)
		}

		apiToken, err := CreateAPIToken()
		if err != nil {
			return fmt.Errorf("Error creating API token: %s", err)
		}

		company := entities.Company{
			Name:     request.Name,
			APIToken: apiToken,
		}

		companyID, err = repositories.CreateCompany(&company, tx)
		if err != nil {
			return fmt.Errorf("Error creating company: %s", err)
		}

		newUser := entities.NewUser(request.Name, request.Username, request.Email, request.Password, "", "", "", "", enumerations.ADMIN, companyID)

		err = repositories.CreateUser(newUser, tx)
		if err != nil {
			return fmt.Errorf("Error creating user: %s", err)
		}

		return nil
	})

	companyResponse := responses.CompanyResponse{
		ID:       companyID,
		Name:     request.Name,
		ApiToken: apiToken,
	}
	return companyResponse
}

func FindCompanies() []responses.CompanyResponse {
	companies, err := repositories.FindCompanies()

	if err != nil {
		return []responses.CompanyResponse{}
	}

	companiesResponse := []responses.CompanyResponse{}

	for _, company := range companies {
		companiesResponse = append(companiesResponse, *MapToCompanyResponse(&company))
	}

	return companiesResponse
}

func FindCompanyByID(id int) *responses.CompanyResponse {
	company, err := repositories.FindCompanyById(id)

	if err != nil {
		exceptions.ThrowNotFoundException(fmt.Sprintf("Company with %d was not found", id))
	}

	return MapToCompanyResponse(company)
}

func CreateAPIToken() (string, error) {
	const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := make([]byte, 10)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes), nil
}

func MapToCompanyResponse(company *entities.Company) (response *responses.CompanyResponse) {
	return &responses.CompanyResponse{
		ID:       company.ID,
		Name:     company.Name,
		ApiToken: company.APIToken,
	}
}
