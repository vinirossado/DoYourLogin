package services

import (
	"doYourLogin/source/domain/entities"
	"doYourLogin/source/infra"
	"doYourLogin/source/middlewares"
	"fmt"
	"time"
)

func CreateLog(logger interface{}, tx *infra.TransactionalOperation) {

	log := &entities.Logs{
		UserID:    middlewares.TokenClaims.ID,
		CompanyID: middlewares.TokenClaims.CompanyID,
		Error:     "Error",
		Message:   fmt.Sprintf("ERROR"),
		Level:     1,
		Route:     "/",
		Source:    "aqui",
		Timestamp: time.Now(),
	}

	_ = repositories.CreateLog(log, tx)

}
