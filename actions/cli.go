package actions

import (
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/domain-based-scaffold/actions/initialization"
	"github.com/lancer-kit/domain-based-scaffold/config"
	"github.com/lancer-kit/domain-based-scaffold/workers"
	"github.com/urfave/cli"
)

const FlagConfig = "config"

func GetFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  FlagConfig + ", c",
			Value: "./config.yaml",
		},
	}
}

func GetCommands() []cli.Command {
	return []cli.Command{
		migrateCommand,
		serveCommand,
	}
}

var serveCommand = cli.Command{
	Name:   "serve",
	Usage:  "starts " + config.ServiceName + " workers",
	Action: serveAction,
}

func serveAction(c *cli.Context) error {
	cfg := initialization.Init(c)
	chief := workers.GetChief(log.Get(), cfg.Workers)
	chief.Run()
	return nil
}
