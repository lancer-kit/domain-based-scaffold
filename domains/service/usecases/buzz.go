package usecases

import (
	"database/sql"

	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/domain-based-scaffold/domains/service/entities"
	"github.com/lancer-kit/domain-based-scaffold/domains/service/repo"
	"github.com/lancer-kit/domain-based-scaffold/errors"
	"github.com/sirupsen/logrus"
)

type BuzzCase struct {
	repo   repo.Repo
	logger *logrus.Entry
}

func NewBuzzCase(repoObj repo.Repo, logger *logrus.Entry) *BuzzCase {
	return &BuzzCase{logger: logger, repo: repoObj}
}

func (buzz *BuzzCase) Add(data *entities.BuzzFeed) (*entities.BuzzFeed, errors.ErrorWithKind) {
	buzz.logger.Debug("Trying to write data into database")

	dataQ := buzz.repo.BuzzFeed()
	err := dataQ.Insert(data)
	if err != nil {
		buzz.logger.WithError(err).Error("can not insert data into database")
		return nil, errors.Communication(err, "can not insert data into database")
	}

	buzz.logger.Debug("Data has been written successfully")
	return data, nil
}

func (buzz *BuzzCase) AllBuzz(pageQuery db.PageQuery) ([]entities.BuzzFeed, errors.ErrorWithKind) {
	dbQuery := buzz.repo.BuzzFeed()

	records, err := dbQuery.SetPage(&pageQuery).Select()
	if err != nil {
		buzz.logger.WithError(err).Error("unable to select")
		if err == sql.ErrNoRows {
			return nil, errors.DataAccess(err, "unable to select")
		}
		return nil, errors.Communication(err, "unable to select")
	}

	buzz.logger.Debug("buzz records was successfully obtained")
	return records, nil
}

func (buzz *BuzzCase) GetBuzz(uid int64) (*entities.BuzzFeed, errors.ErrorWithKind) {
	dataQ := buzz.repo.BuzzFeed()

	res, err := dataQ.GetByID(uid)
	if err != nil {
		buzz.logger.WithError(err).Error("can not get by id")
		if err == sql.ErrNoRows {
			return nil, errors.DataAccess(err, "can not get by id")
		}
		return nil, errors.Communication(err, "can not get by id")
	}

	return res, nil
}

type ChangeBuzzData struct {
	Description string `json:"description"`
}

func (buzz *BuzzCase) ChangeBuzz(uid int64, data *ChangeBuzzData) errors.ErrorWithKind {
	dataQ := buzz.repo.BuzzFeed()
	err := dataQ.UpdateBuzzDescription(uid, data.Description)
	if err != nil {
		return errors.Communication(err, "can not update BuzzFeed description")
	}

	buzz.logger.Debug("Data has been written successfully")
	return nil
}

func (buzz *BuzzCase) DeleteBuzz(uid int64) errors.ErrorWithKind {
	dataQ := buzz.repo.BuzzFeed()
	err := dataQ.DeleteByID(uid)
	if err != nil {
		buzz.logger.WithError(err).Error("can not delete BuzzFeed")
		return errors.Communication(err, "can not delete BuzzFeed")
	}

	buzz.logger.Debug("Data has been deleted successfully")
	return nil
}
