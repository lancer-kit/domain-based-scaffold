package workers

import (
	"github.com/lancer-kit/domain-based-scaffold/config"
	"github.com/lancer-kit/uwe/v2"
	"github.com/sirupsen/logrus"
)

func GetChief(logger *logrus.Entry, workers []uwe.WorkerName) uwe.Chief {
	var wMap = map[uwe.WorkerName]uwe.Worker{
		config.WorkerAPIServer: Server(logger.WithField("worker", config.WorkerFooBar), config.Config()),
		config.WorkerFooBar:    NewWorker(logger.WithField("worker", config.WorkerFooBar)),
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

		logger.WithFields(event.Fields).
			Log(level, event.Message)

	})

	return chief
}
