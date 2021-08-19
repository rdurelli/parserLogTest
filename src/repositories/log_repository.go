package repositories

import (
	"fmt"
	"log"
	"strings"

	"github.com/rdurelli/loggingParser/src/lib"

	"github.com/rdurelli/loggingParser/src/models"
)

type logRepository struct {
	Db lib.Db
}

func NewLogReposiry(Db lib.Db) Repository {
	return &logRepository{
		Db: Db,
	}
}

func (lR logRepository) Save(lo *models.Log) (*models.Log, error) {
	stmt, err := lR.Db.Db.Prepare("INSERT INTO log (date, time, http_code, http_method, path) VALUES (?, ?, ?, ?, ?)")
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
	_, err := lR.Db.Db.Exec(stmt, valueArgs...)
	return err
}

func (lR logRepository) Find(offset string, limit string) (*[]models.Log, error) {
	var logs = make([]models.Log, 0)
	stmt, err := lR.Db.Db.Prepare("SELECT ID, date, time, http_code, http_method, path FROM log LIMIT ?  OFFSET ?")
	if err != nil {
		log.Fatal("Not possible Create statement to save LOG", err)
		return nil, err
	}
	rows, err := stmt.Query(limit, offset)
	defer rows.Close()
	if err != nil {
		log.Fatal("Not possible to save into LOG", err)
		return nil, err
	}
	for rows.Next() {
		var log models.Log
		if err := rows.Scan(&log.ID, &log.Date, &log.Time, &log.HttpCode,
			&log.HttpMethod, &log.Path); err != nil {
			return &logs, err
		}
		logs = append(logs, log)
	}
	if err = rows.Err(); err != nil {
		return &logs, err
	}
	return &logs, nil
}
