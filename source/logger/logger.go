package trace_logger

import (
	"doYourLogin/source/domain/entities"
	"doYourLogin/source/middlewares"
	"doYourLogin/source/repositories"
	"time"
)

func insert(logger *entities.Logs, tx *repositories.TransactionalOperation) {
	_ = repositories.CreateLog(logger, tx)
}

func BuildLogger(message, route, source string, error uint, tx *repositories.TransactionalOperation) {

	logger := entities.Logs{
		UserID:    middlewares.TokenClaims.ID,
		CompanyID: middlewares.TokenClaims.CompanyID,
		Error:     error,
		Message:   message,
		//Level:     level,
		Route:     route,
		Source:    source,
		Timestamp: time.Now(),
	}
	insert(&logger, tx)
}

//
//func (l *entities.Logs) FlushCacheToDatabase() {}
//
//func handleFlowCompletion(log *Logs) {}
