package services

import (
	"crypto/rand"
	"doYourLogin/source/domain/entities"
	"doYourLogin/source/domain/exceptions"
	"doYourLogin/source/domain/requests"
	"doYourLogin/source/domain/responses"
	"doYourLogin/source/repositories"
	"fmt"
)

func CreateCompany(request *requests.CompanyRequest) responses.CompanyResponse {
	var Id uint
	var APIToken string

	repositories.UsingTransactional(func(tx *repositories.TransactionalOperation) error {
		exists := repositories.ExistsUserByUsername(request.Name)

		if exists {
			return exceptions.BadRequestException(
				fmt.Sprintf("Company %s already exists", request.Name),
			)
		}

		apiToken := CreateAPIToken()

		company := entities.Company{
			Name:     request.Name,
			APIToken: apiToken,
		}

		id, err := repositories.CreateCompany(&company, tx)

		if err != nil {
			return exceptions.InternalServerException(
				fmt.Sprintf("Error while trying to insert new User with error: %s", err),
			)
		}

		Id = id
		APIToken = apiToken

		return nil
	})

	companyResponse := responses.CompanyResponse{
		ID:       Id,
		Name:     request.Name,
		ApiToken: APIToken,
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

func CreateAPIToken() string {
	const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := make([]byte, 10)

	if _, err := rand.Read(bytes); err != nil {
		return ""
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes)
}

func MapToCompanyResponse(company *entities.Company) (response *responses.CompanyResponse) {
	return &responses.CompanyResponse{
		ID:       company.ID,
		Name:     company.Name,
		ApiToken: company.APIToken,
	}
}
