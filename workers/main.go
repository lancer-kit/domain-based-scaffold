package workers

import (
	"github.com/lancer-kit/service-scaffold/config"
	"github.com/lancer-kit/service-scaffold/workers/api"
	"github.com/lancer-kit/service-scaffold/workers/foobar"
	"github.com/lancer-kit/uwe"
	"github.com/lancer-kit/armory/log"
)

var WorkerChief uwe.Chief

func GetChief() *uwe.Chief {
	WorkerChief = uwe.Chief{EnableByDefault: true}
	WorkerChief.AddWorker(config.WorkerAPIServer, api.Server())
	WorkerChief.AddWorker(config.WorkerFooBar, &foobar.Worker{})

	WorkerChief.Init(log.Default)

	return &WorkerChief
}
