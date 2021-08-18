package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/rdurelli/loggingParser/src/models"
)

type logRepository struct {
	Db *sql.DB
}

func NewLogReposiry(db *sql.DB) Repository {
	return &logRepository{
		Db: db,
	}
}

func (lR logRepository) Save(lo *models.Log) (*models.Log, error) {
	stmt, err := lR.Db.Prepare("INSERT INTO log (date, time, http_code, http_method, path) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Not possible Create statement to save LOG", err)
		return nil, err
	}
	result, err := stmt.Exec(lo.Date, lo.Time, lo.HttpCode, lo.HttpMethod, lo.Path)
	if err != nil {
		log.Fatal("Not possible to save into LOG", err)
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Unable to get the ID", err)
		return nil, err
	}
	lo.ID = uint(id)
	return lo, nil
}

func (lR logRepository) SaveBatch(logs *[]models.Log) error {
	valueStrings := make([]string, 0, len(*logs))
	valueArgs := make([]interface{}, 0, len(*logs)*5)
	for _, log := range *logs {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, log.Date)
		valueArgs = append(valueArgs, log.Time)
		valueArgs = append(valueArgs, log.HttpCode)
		valueArgs = append(valueArgs, log.HttpMethod)
		valueArgs = append(valueArgs, log.Path)
	}
	stmt := fmt.Sprintf("INSERT INTO log (date, time, http_code, http_method, path) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err := lR.Db.Exec(stmt, valueArgs...)
	return err
}

func (lR logRepository) Find(offset string, limit string) (*[]models.Log, error) {
	panic("implement me")
}