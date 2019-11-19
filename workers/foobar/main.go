package foobar

import (
	"time"

	"github.com/lancer-kit/uwe/v2"
	"github.com/sirupsen/logrus"
)

type Worker struct {
	logger *logrus.Entry
	init   bool
}

func NewWorker(logger *logrus.Entry) uwe.Worker {
	return &Worker{
		logger: logger,
	}
}

func (d *Worker) Init() error {
	d.init = true
	return nil
}

func (d Worker) Run(wCtx uwe.Context) error {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			d.logger.Info("Perform my task")
		case <-wCtx.Done():
			ticker.Stop()
			d.logger.Info("Receive exit code, stop working")
			return nil
		}
	}

}
