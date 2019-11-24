package delivery

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/lancer-kit/armory/api/httpx"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/domain-based-scaffold/domains/service/entities"
	"github.com/lancer-kit/domain-based-scaffold/domains/service/usecases"
	"github.com/pkg/errors"
)

func (h Handlers) AddDocument(w http.ResponseWriter, r *http.Request) {
	logger := log.IncludeRequest(h.logger, r)

	data := new(entities.CustomDocument)
	err := httpx.ParseJSONBody(r, data)
	if err != nil {
		logger.WithError(err).Error("can not parse the body")
		render.BadRequest(w, "invalid body, must be json")
		return
	}

	logger.Debug("Trying to write data into couchdb")

	err = usecases.NewDocumentCases(h.repo, logger).AddDocument(data)
	if err != nil {
		logger.WithError(err).Error("Can not insert data into database")
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	logger.Debug("Data has been written successfully")
	render.WriteJSON(w, 201, data)
}

func (h Handlers) GetAllDocument(w http.ResponseWriter, r *http.Request) {
	uid, err := extractUID(r, "id")
	if err != nil {
		log.IncludeRequest(h.logger, r).WithError(err).Error("can not parse id")
		render.BadRequest(w, "invalid id, should be a number")
		return
	}

	logger := log.IncludeRequest(h.logger, r).WithField("query_id", uid)

	pageQuery, err := db.ParsePageQuery(r.URL.Query())
	if err != nil {
		logger.WithError(err).Error("invalid page query")
		render.BadRequest(w, "invalid page query")
		return
	}

	res, err := usecases.NewDocumentCases(h.repo, logger).GetAllDocument(pageQuery)
	if err != nil {
		logger.WithError(err).Error("can not to get documents")
		render.ServerError(w)
		return
	}

	render.RenderListWithPages(w, pageQuery, int64(len(res)), res)
}

func (h Handlers) GetDocument(w http.ResponseWriter, r *http.Request) {
	uid, err := extractUID(r, "id")
	if err != nil {
		log.IncludeRequest(h.logger, r).WithError(err).Error("can not parse id")
		render.BadRequest(w, "invalid id, should be a number")
		return
	}

	logger := log.DefaultForRequest(r).WithField("query_uid", uid)

	res, err := usecases.NewDocumentCases(h.repo, logger).GetDocuments(uid)
	if err != nil {
		logger.WithError(err).Error("can not to get document")
		render.ServerError(w)
		return
	}

	render.Success(w, res)
}

func (h Handlers) ChangeDocument(w http.ResponseWriter, r *http.Request) {
	uid, err := extractUID(r, "id")
	if err != nil {
		log.IncludeRequest(h.logger, r).WithError(err).Error("can not parse id")
		render.BadRequest(w, "invalid id, should be a number")
		return
	}

	logger := log.IncludeRequest(h.logger, r).WithField("query_id", uid)

	data := new(entities.CustomDocument)
	if err = httpx.ParseJSONBody(r, data); err != nil {
		logger.WithError(err).Error("can not parse the body")
		render.BadRequest(w, "invalid body, must be json")
		return
	}

	caseError := usecases.NewDocumentCases(h.repo, logger).ChangeDocument(uid, data)
	if caseError != nil {
		logger.WithError(err).Error("сan not update document")
		render.ServerError(w)
		return
	}

	logger.Debug("Data has been written successfully")
	render.Success(w, "Document was updated successful")
}

func (h Handlers) DeleteDocument(w http.ResponseWriter, r *http.Request) {
	uid, err := extractUID(r, "id")
	if err != nil {
		log.IncludeRequest(h.logger, r).WithError(err).Error("can not parse id")
		render.BadRequest(w, "invalid id, should be a number")
		return
	}

	logger := log.IncludeRequest(h.logger, r).WithField("query_uid", uid)

	err = usecases.NewDocumentCases(h.repo, logger).DeleteDocument(uid)
	if err != nil {
		logger.WithError(err).Error("can not delete document")
		render.ServerError(w)
		return
	}

	render.Success(w, "Document was successfully deleted")
}

func extractUID(r *http.Request, queryKey string) (int64, error) {
	uid := chi.URLParam(r, queryKey)

	idINT, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return idINT, errors.Wrap(err, "invalid id, should be a number")
	}
	return idINT, nil

}
