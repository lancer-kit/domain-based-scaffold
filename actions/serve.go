package actions

import (
	"github.com/lancer-kit/domain-based-scaffold/actions/initialization"
	"github.com/lancer-kit/domain-based-scaffold/config"
	"github.com/lancer-kit/domain-based-scaffold/workers"
	"github.com/urfave/cli"
)

var serveCommand = cli.Command{
	Name:   "serve",
	Usage:  "starts " + config.ServiceName + " workers",
	Action: serveAction,
}

func serveAction(c *cli.Context) error {
	cfg := initialization.Init(c)

	err := workers.GetChief().Run(cfg.Workers...)
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	return nil
}
