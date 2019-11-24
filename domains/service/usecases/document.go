package usecases

import (
	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/domain-based-scaffold/domains/service/entities"
	"github.com/lancer-kit/domain-based-scaffold/domains/service/repo"
	"github.com/lancer-kit/domain-based-scaffold/errors"
	"github.com/sirupsen/logrus"
)

type DocumentCases struct {
	logger *logrus.Entry
	repo   repo.Repo
}

func NewDocumentCases(repoObj repo.Repo, logger *logrus.Entry) *DocumentCases {
	return &DocumentCases{repo: repoObj, logger: logger}
}

func (document *DocumentCases) AddDocument(data *entities.CustomDocument) errors.ErrorWithKind {
	document.logger.Debug("Trying to write data into couchdb")
	docQ, err := document.repo.CustomDocument()
	if err != nil {
		document.logger.WithError(err).Error("can not create documents connector")
		return errors.Communication(err, "can not create documents connector")
	}

	err = docQ.AddDocument(data)
	if err != nil {
		document.logger.WithError(err).Error("Can not insert data into database")
		return errors.DataAccess(err, "can not insert data into database")
	}

	document.logger.Debug("Data has been written successfully")
	return nil
}

func (document *DocumentCases) GetAllDocument(pageQuery db.PageQuery) ([]entities.CustomDocument, errors.ErrorWithKind) {
	docQ, err := document.repo.CustomDocument()
	if err != nil {
		document.logger.WithError(err).Error("can not create documents connector")
		return nil, errors.Communication(err, "can not create documents connector")
	}

	res, err := docQ.GetAllDocument(pageQuery)
	if err != nil {
		document.logger.WithError(err).Error("can not to get documents")
		return nil, errors.DataAccess(err, "can not to get documents")
	}

	return res, nil
}

func (document *DocumentCases) GetDocuments(uid int64) ([]entities.CustomDocument, errors.ErrorWithKind) {
	docQ, err := document.repo.CustomDocument()
	if err != nil {
		document.logger.WithError(err).Error("can not create documents connector")
		return nil, errors.Communication(err, "can not create documents connector")
	}

	res, err := docQ.GetDocument(uid)
	if err != nil {
		document.logger.WithError(err).Error("can not to get document")
		return nil, errors.DataAccess(err, "can not to get document")
	}

	return res, nil
}

func (document *DocumentCases) ChangeDocument(uid int64, data *entities.CustomDocument) errors.ErrorWithKind {
	docQ, err := document.repo.CustomDocument()
	if err != nil {
		document.logger.WithError(err).Error("can not create documents connector")
		return errors.Communication(err, "can not create documents connector")
	}

	err = docQ.UpdateDocument(uid, data)
	if err != nil {
		document.logger.WithError(err).Error("сan not update document")
		return errors.DataAccess(err, "can not create documents connector")
	}

	document.logger.Debug("Data has been written successfully")
	return nil
}

func (document *DocumentCases) DeleteDocument(uid int64) errors.ErrorWithKind {
	docQ, err := document.repo.CustomDocument()
	if err != nil {
		document.logger.WithError(err).Error("can not create documents connector")
		return errors.Communication(err, "can not create documents connector")
	}

	err = docQ.DeleteDocument(uid)
	if err != nil {
		document.logger.WithError(err).Error("can not delete document")
		return errors.DataAccess(err, "can not create custom document")
	}

	return nil
}
