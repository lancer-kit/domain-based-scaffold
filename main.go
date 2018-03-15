package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/inn4sc/vcg-go-common/log"
	"github.com/inn4sc/vcg-go-common/routines"
	"github.com/inn4sc/go-skeleton/services/api"
	"github.com/inn4sc/go-skeleton/services/foobar"
	"github.com/urfave/cli"
)

var WorkerChief routines.Chief

func init() {
	WorkerChief = routines.Chief{}
	WorkerChief.AddWorker("api-server", &api.Server{})
	WorkerChief.AddWorker("foobar", &foobar.Service{})
}

var EnabledServices = []string{
	"alpha-worker",
	"beta-worker",
}

func main() {
	cmd := cli.NewApp()
	cmd.Usage = "A Skeleton service"
	cmd.Version = "0.4.2"

	cmd.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "starts skeleton service",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config, c",
					Value: "./config.yaml",
				},
			},
			Action: runAction,
		},
	}
}

func runAction(c *cli.Context) error {
	done := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	for _, serviceName := range EnabledServices {
		WorkerChief.EnableWorker(serviceName)
	}

	WorkerChief.InitWorkers(log.Default)
	go func() {
		defer close(done)
		WorkerChief.Start(ctx)
	}()

	log.Default.Info("Payment Gate started")

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
}