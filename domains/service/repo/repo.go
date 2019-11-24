package repo

import (
	"github.com/lancer-kit/armory/db"
	"github.com/sirupsen/logrus"
)

// Repo is a top level interface for interaction with database.
type Repo interface {
	db.Transactional
	Clone() Repo

	BuzzFeed() BuzzFeedQI
	CustomDocument() (CustomDocumentRepo, error)
}

// repo implementation of the `Repo` interface.
type repo struct {
	*db.SQLConn
}

func (q repo) Clone() Repo {
	return &repo{
		SQLConn: q.SQLConn.Clone(),
	}
}

// NewRepo returns initialized instance of the `Repo`.
func NewRepo(config db.Config, logger *logrus.Entry) (Repo, error) {
	sqlConn, err := db.NewConnector(config, logger)
	if err != nil {
		return nil, err
	}
	return &repo{
		SQLConn: sqlConn,
	}, nil
}
