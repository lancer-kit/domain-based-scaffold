package models

import (
	cdb "github.com/leesper/couchdb-golang"
	"github.com/pkg/errors"
	"gitlab.inn4science.com/gophers/service-kit/db"
	"gitlab.inn4science.com/gophers/service-scaffold/config"
)

type CustomDocument struct {
	Id         int64  `json:"id"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	cdb.Document
}

type customDocumentQ struct {
	dbInstance *cdb.Database
}

func CreateCustomDocumentQ() (*customDocumentQ, error) {
	newDocInstance := new(customDocumentQ)

	dbInstance, err := cdb.NewDatabase(config.Config().CouchDB)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to connect to couchdb")
	}

	newDocInstance.dbInstance = dbInstance
	return newDocInstance, nil
}

func (d *customDocumentQ) AddDocument(doc *CustomDocument) error {
	err := cdb.Store(d.dbInstance, doc)
	if err != nil {
		return errors.Wrap(err, "Unable to write into couchdb")
	}

	return nil
}

func (d *customDocumentQ) GetAllDocument(pQ db.PageQuery) ([]CustomDocument, error) {
	fields := []string{"id", "firstName", "secondName"}
	res, err := d.dbInstance.Query(fields, `exists(id,true)`, nil, nil, nil, nil)

	if err != nil {
		return nil, errors.Wrap(err, "Unable to write into couchdb")
	}

	resSlice := make([]CustomDocument, 0)
	for _, v := range res {
		obj := new(CustomDocument)
		cdb.FromJSONCompatibleMap(obj, v)
		resSlice = append(resSlice, *obj)
	}

	return resSlice, nil
}

func (d *customDocumentQ) GetDocument() {

}

func (d *customDocumentQ) DeleteDocument() {

}
