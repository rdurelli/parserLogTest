package services

import (
	"github.com/rdurelli/loggingParser/src/lib"
	"github.com/rdurelli/loggingParser/src/models"
	"github.com/rdurelli/loggingParser/src/repositories"
)

type LogService struct {
	logger     lib.Logger
	repository repositories.Repository
}

func NewLogService(logger lib.Logger, repository repositories.Repository) LogService {
	return LogService{
		logger:     logger,
		repository: repository,
	}
}

func (lS LogService) Find(offset string, limit string) (*[]models.Log, error) {
	panic("implement me")
}
