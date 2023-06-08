package services

import (
	"crypto/rand"
	"doYourLogin/source/domain/entities"
	"doYourLogin/source/domain/enumerations"
	"doYourLogin/source/domain/exceptions"
	"doYourLogin/source/domain/requests"
	"doYourLogin/source/domain/responses"
	"doYourLogin/source/middlewares"
	"doYourLogin/source/repositories"
	"doYourLogin/source/utils"
	"encoding/json"
	"fmt"
	"time"
)

func CreateCompany(request *requests.CompanyRequest) responses.CompanyResponse {
	var companyID uint
	var apiToken string

	repositories.UsingTransactional(func(tx *repositories.TransactionalOperation) error {
		err := tx.BeginTransaction()
		if err != nil {
			return err
		}

		if repositories.ExistsUserByUsername(request.Name) {
			return fmt.Errorf("company %s already exists", request.Name)
		}

		apiToken, err := CreateAPIToken()
		if err != nil {
			return fmt.Errorf("error creating API token: %s", err)
		}

		company := entities.Company{
			Name:            request.Name,
			APIToken:        apiToken,
			ExpirationToken: time.Now().Add(time.Hour),
		}

		companyID, err = repositories.CreateCompany(&company, tx)
		if err != nil {
			return fmt.Errorf("error creating company: %s", err)
		}

		newUser := entities.NewUser(request.Name, request.Username, request.Email, request.Password, "", "", "", "", enumerations.ADMIN, companyID)

		err = repositories.CreateUser(newUser, tx)
		if err != nil {
			return fmt.Errorf("error creating user: %s", err)
		}

		return nil
	})

	companyResponse := responses.CompanyResponse{
		ID:       companyID,
		Name:     request.Name,
		APIToken: apiToken,
	}

	return companyResponse
}

func FindCompanies() []responses.CompanyResponse {
	if utils.Cache == nil {
		fmt.Println("Cache não inicializado.")
		return []responses.CompanyResponse{}
	}

	if entry, err := utils.Cache.Get("companies"); err == nil {
		var companiesResponse []responses.CompanyResponse
		err = json.Unmarshal(entry, &companiesResponse)
		if err == nil {
			return companiesResponse
		}
	}

	companies, err := repositories.FindCompanies()
	if err != nil {
		return []responses.CompanyResponse{}
	}

	var companiesResponse []responses.CompanyResponse
	utils.Map(&companies, &companiesResponse)

	responseBytes, err := json.Marshal(companiesResponse)
	if err == nil {
		err = utils.Cache.Set("companies", responseBytes)
		if err != nil {
			fmt.Println("Erro ao armazenar no cache:", err)
		}
	}
	return companiesResponse
}

func FindMyCompany() *responses.CompanyResponse {
	companyID := middlewares.TokenClaims.CompanyID
	company, err := repositories.FindCompanyById(companyID)

	if err != nil {
		exceptions.ThrowNotFoundException(fmt.Sprintf("Company with %d was not found", companyID))
	}

	var companyResponse responses.CompanyResponse

	utils.Map(company, &companyResponse)

	return &companyResponse
}

func ActivateAccount(apiToken string) *responses.CompanyResponse {
	repositories.UsingTransactional(func(tx *repositories.TransactionalOperation) error {
		company, _ := repositories.FindCompanyByApiToken(apiToken, tx)
		company.Active = true

		_ = repositories.ActivateAccount(company, tx)

		return nil
	})
	return nil
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
