package entities

import (
	cdb "github.com/leesper/couchdb-golang"
)

type CustomDocument struct {
	Id         int64  `json:"id"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	cdb.Document
}
