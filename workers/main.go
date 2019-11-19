package workers

import (
	"github.com/lancer-kit/domain-based-scaffold/config"
	"github.com/lancer-kit/domain-based-scaffold/workers/api"
	"github.com/lancer-kit/domain-based-scaffold/workers/foobar"
	"github.com/lancer-kit/uwe/v2"
	"github.com/sirupsen/logrus"
)

func GetChief(logger *logrus.Entry, workers []uwe.WorkerName) uwe.Chief {
	var wMap = map[uwe.WorkerName]uwe.Worker{
		config.WorkerAPIServer: api.Server(logger.WithField("worker", config.WorkerFooBar), config.Config().Api),
		config.WorkerFooBar:    foobar.NewWorker(logger.WithField("worker", config.WorkerFooBar)),
	}

	chief := uwe.NewChief()
	for _, wName := range workers {
		chief.AddWorker(wName, wMap[wName])
	}

	chief.UseDefaultRecover()
	chief.SetEventHandler(func(event uwe.Event) {
		var level logrus.Level
		switch event.Level {
		case uwe.LvlFatal, uwe.LvlError:
			level = logrus.ErrorLevel
		case uwe.LvlInfo:
			level = logrus.InfoLevel
		}
		logger.WithFields(event.Fields).Log(level, event.Message)

	})

	return chief
}
