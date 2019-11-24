package delivery

import (
	"net/http"

	"github.com/lancer-kit/armory/api/httpx"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/domain-based-scaffold/domains/service/entities"
	"github.com/lancer-kit/domain-based-scaffold/domains/service/repo"
	"github.com/lancer-kit/domain-based-scaffold/domains/service/usecases"
	"github.com/sirupsen/logrus"
)

type Handlers struct {
	repo   repo.Repo
	logger *logrus.Entry
}

func NewHandlers(repoObj repo.Repo, logger *logrus.Entry) Handlers {
	return Handlers{
		repo:   repoObj,
		logger: logger,
	}
}
func (h Handlers) AddBuzz(w http.ResponseWriter, r *http.Request) {
	logger := log.IncludeRequest(h.logger, r)

	data := new(entities.BuzzFeed)
	err := httpx.ParseJSONBody(r, data)
	if err != nil {
		logger.WithError(err).Error("can not parse the body")
		render.BadRequest(w, "invalid body, must be json")
		return
	}

	data, caseErr := usecases.NewBuzzCase(h.repo.Clone(), logger).Add(data)
	if caseErr != nil {
		logger.WithError(err).Error("Can not insert data into database")
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	logger.Debug("Data has been written successfully")
	render.WriteJSON(w, 201, data)
}

func (h Handlers) AllBuzz(w http.ResponseWriter, r *http.Request) {
	logger := log.IncludeRequest(h.logger, r)

	pageQuery, err := db.ParsePageQuery(r.URL.Query())
	if err != nil {
		logger.WithError(err).Error("invalid page query")
		render.BadRequest(w, "invalid page query")
		return
	}

	ols, err := usecases.NewBuzzCase(h.repo.Clone(), logger).AllBuzz(pageQuery)
	if err != nil {
		logger.WithError(err).Error("unable to select")
		render.ServerError(w)
		return
	}

	logger.Debug("Buzz instances was successfully obtained")
	render.RenderListWithPages(w, pageQuery, int64(len(ols)), ols)
}

func (h Handlers) GetBuzz(w http.ResponseWriter, r *http.Request) {
	uid, err := extractUID(r, "id")
	if err != nil {
		log.IncludeRequest(h.logger, r).WithError(err).Error("can not parse id")
		render.BadRequest(w, "invalid id, should be a number")
		return
	}

	logger := log.IncludeRequest(h.logger, r).WithField("query_uid", uid)

	res, err := usecases.NewBuzzCase(h.repo.Clone(), logger).GetBuzz(uid)
	if err != nil {
		logger.WithError(err).Error("can not get by id")
		render.ServerError(w)
		return
	}

	render.Success(w, res)
}

func (h Handlers) ChangeBuzz(w http.ResponseWriter, r *http.Request) {
	uid, err := extractUID(r, "id")
	if err != nil {
		log.IncludeRequest(h.logger, r).WithError(err).Error("can not parse id")
		render.BadRequest(w, "invalid id, should be a number")
		return
	}

	logger := log.IncludeRequest(h.logger, r).WithField("query_id", uid)

	data := new(usecases.ChangeBuzzData)
	err = httpx.ParseJSONBody(r, data)
	if err != nil {
		logger.WithError(err).Error("can not parse the body")
		render.BadRequest(w, "invalid body, must be json")
		return
	}

	err = usecases.NewBuzzCase(h.repo.Clone(), logger).ChangeBuzz(uid, data)
	if err != nil {
		logger.WithError(err).Error("Can not update buzzfeed description")
		render.ServerError(w)
		return
	}

	logger.Debug("Data has been written successfully")
	render.Success(w, data)
}

func (h Handlers) DeleteBuzz(w http.ResponseWriter, r *http.Request) {
	uid, err := extractUID(r, "id")
	if err != nil {
		log.IncludeRequest(h.logger, r).WithError(err).Error("can not parse id")
		render.BadRequest(w, "invalid id, should be a number")
		return
	}

	logger := log.IncludeRequest(h.logger, r).WithField("query_uid", uid)

	err = usecases.NewBuzzCase(h.repo.Clone(), logger).DeleteBuzz(uid)
	if err != nil {
		logger.WithError(err).Error("can not delete BuzzFeed")
		render.ServerError(w)
		return
	}

	logger.Debug("Data has been deleted successfully")
	render.Success(w, "success")
}
