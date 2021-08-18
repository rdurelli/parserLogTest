package repositories

import "github.com/rdurelli/loggingParser/src/models"

type Reader interface {
	Find(offset string, limit string) (*[]models.Log, error)
}

type Writer interface {
	Save(log *models.Log) (*models.Log, error)
	SaveBatch(logs *[]models.Log) error
}

type Repository interface {
	Writer
	Reader
}
