package models

import "fmt"

type Log struct {
	ID         uint   `json:"id,omitempty"`
	Date       string `json:"date,omitempty"`
	Time       string `json:"time,omitempty"`
	HttpCode   string `json:"http_code,omitempty"`
	HttpMethod string `json:"http_method,omitempty"`
	Path       string `json:"path,omitempty"`
}

func (l Log) String() string {
	return fmt.Sprintf("[%s, %s, %s, %s, %s]", l.Date, l.Time, l.HttpCode, l.HttpMethod, l.Path)
}
