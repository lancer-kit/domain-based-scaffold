package main

import (
	"log"
	"os"

	"github.com/lancer-kit/domain-based-scaffold/actions"
	"github.com/lancer-kit/domain-based-scaffold/config"
	"github.com/lancer-kit/domain-based-scaffold/info"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "A " + config.ServiceName + " service"
	app.Version = info.App.Version
	app.Flags = actions.GetFlags()
	app.Commands = actions.GetCommands()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
