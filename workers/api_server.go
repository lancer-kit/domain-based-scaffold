package workers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/auth"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/domain-based-scaffold/config"
	"github.com/lancer-kit/domain-based-scaffold/domains/service/delivery"
	"github.com/lancer-kit/domain-based-scaffold/domains/service/repo"
	"github.com/lancer-kit/domain-based-scaffold/info"
	"github.com/lancer-kit/uwe/v2/presets/api"
	"github.com/sirupsen/logrus"
)

func Server(entry *logrus.Entry, cfg *config.Configuration) *api.Server {
	return api.NewServer(cfg.Api, GetRouter(entry, cfg))
}

func GetRouter(logger *logrus.Entry, config *config.Configuration) http.Handler {
	r := chi.NewRouter()
	repoObj, err := repo.NewRepo(config.DB, logger)
	if err != nil {
		logger.WithError(err).Fatal("unable to initialize repo")
	}

	handlers := delivery.NewHandlers(repoObj, logger)

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(log.NewRequestLogger(logger.Logger))

	if config.Api.EnableCORS {
		r.Use(getCORS().Handler)
	}

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	if config.Api.ApiRequestTimeout > 0 {
		t := time.Duration(config.Api.ApiRequestTimeout)
		r.Use(middleware.Timeout(t * time.Second))
	}

	r.Route("/dev", func(r chi.Router) {
		r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
			render.Success(w, info.App)
		})

		r.Route("/", func(r chi.Router) {
			r.Use(auth.ExtractUserID())

			r.Route("/{mId}/buzz", func(r chi.Router) {
				//custom middleware example
				r.Use(delivery.VerifySomethingMiddleware())
				r.Post("/", handlers.AddBuzz)
				r.Get("/", handlers.AllBuzz)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", handlers.GetBuzz)
					r.Put("/", handlers.ChangeBuzz)
					r.Delete("/", handlers.DeleteBuzz)
				})

			})
		})

		r.Route("/couch", func(r chi.Router) {
			r.Post("/", handlers.AddDocument)
			r.Get("/", handlers.GetAllDocument)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetDocument)
				r.Put("/", handlers.ChangeDocument)
				r.Delete("/", handlers.DeleteDocument)
			})
		})
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.ResultNotFound.Render(w)
	})

	return r
}

func getCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "jwt", "X-UID"},
		ExposedHeaders:   []string{"Link", "Content-Length"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
}
