package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"fmt"

	"github.com/urfave/cli"
	"gitlab.inn4science.com/vcg/go-common/log"

	"gitlab.inn4science.com/vcg/go-skeleton/config"
	"gitlab.inn4science.com/vcg/go-skeleton/dbschema"
	"gitlab.inn4science.com/vcg/go-skeleton/workers"
)

func serveAction(c *cli.Context) error {
	initConfig(c.String("config"))
	cfg := config.Config()

	if cfg.AutoMigrate {
		count, err := dbschema.Migrate(config.Config().DB, "up")
		if err != nil {
			log.Default.WithError(err).Error("Migrations failed")
		}
		log.Default.Info(fmt.Sprintf("Applied %d %s migration", count, "up"))
	}

	done := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	for _, workerName := range cfg.Workers {
		workers.WorkerChief.EnableWorker(workerName)
	}

	workers.WorkerChief.InitWorkers(log.Default)
	go func() {
		defer close(done)
		workers.WorkerChief.Start(ctx)
	}()

	log.Default.Info(config.ServiceName + " started")

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)

	exitSignal := <-gracefulStop
	log.Default.WithField("signal", exitSignal).
		Info("Received signal. Terminating service...")
	cancel()

	select {
	case <-done:
		log.Default.Info("Graceful exit.")
		return nil
	case <-time.NewTimer(60 * time.Second).C:
		log.Default.Warn("Graceful exit timeout exceeded. Force exit.")
		return cli.NewExitError("Graceful exit timeout exceeded", 1)
	}

	return nil
}

var serveCommand = cli.Command{

	Name:   "serve",
	Usage:  "starts " + config.ServiceName + " workers",
	Flags:  cfgFlag,
	Action: serveAction,
}
