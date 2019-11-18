package workers

import (
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/domain-based-scaffold/config"
	"github.com/lancer-kit/domain-based-scaffold/workers/api"
	"github.com/lancer-kit/domain-based-scaffold/workers/foobar"
	"github.com/lancer-kit/uwe"
)

var WorkerChief uwe.Chief

func GetChief() *uwe.Chief {
	WorkerChief = uwe.Chief{EnableByDefault: true}
	WorkerChief.AddWorker(config.WorkerAPIServer, api.Server())
	WorkerChief.AddWorker(config.WorkerFooBar, &foobar.Worker{})

	WorkerChief.Init(log.Default)

	return &WorkerChief
}
