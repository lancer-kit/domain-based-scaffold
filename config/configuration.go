package config

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/armory/natsx"
	"github.com/lancer-kit/uwe/v2"
	"github.com/lancer-kit/uwe/v2/presets/api"
)

// Configuration main structure of the app configuration.
type Configuration struct {
	Api     api.Config   `json:"api" yaml:"api"`
	Log     log.Config   `json:"log" yaml:"log"`
	DB      db.Config    `json:"db" yaml:"db"` // DB is a database connection parameters.
	NATS    natsx.Config `json:"nats" yaml:"nats"`
	CouchDB string       `json:"couchdb" yaml:"couchdb"` // CouchDB is a couchdb url connection string.

	ModulesInitTimeout int `json:"modules_init_timeout" yaml:"modules_init_timeout"`

	// Workers is a list of workers
	// that must be started, start all if empty.
	Workers []string `json:"workers" yaml:"workers"`
}

func (cfg Configuration) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.DB, validation.Required),
		validation.Field(&cfg.ModulesInitTimeout, validation.Required),
		validation.Field(&cfg.CouchDB, validation.Required),
		validation.Field(&cfg.Api, validation.Required),
		//validation.Field(&cfg.NATS, validation.Required),
		validation.Field(&cfg.Workers, &uwe.WorkerExistRule{
			AvailableWorkers: availableWorkers,
		}),
	)
}

func (cfg Configuration) fillDefaultWorkers() {
	for k := range availableWorkers {
		cfg.Workers = append(cfg.Workers, string(k))
	}
}

func (cfg Configuration) WorkersList() []uwe.WorkerName {
	if config.Workers == nil {
		config.fillDefaultWorkers()
	}
	var list []uwe.WorkerName

	for _, name := range cfg.Workers {
		list = append(list, uwe.WorkerName(name))
	}
	return list
}
