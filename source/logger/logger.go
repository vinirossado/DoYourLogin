package trace_logger

import (
	"doYourLogin/source/domain/entities"
	"doYourLogin/source/middlewares"
	"doYourLogin/source/repositories"
	"fmt"
	"time"
)

func Create(logger interface{}, tx *repositories.TransactionalOperation) {

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
