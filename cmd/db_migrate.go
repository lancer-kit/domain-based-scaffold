package cmd

import (
	"fmt"

	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/service-scaffold/config"
	"github.com/lancer-kit/service-scaffold/dbschema"
	"github.com/urfave/cli"
)

var migrateCommand = cli.Command{
	Name:  "migrate",
	Usage: "apply db migration",

	Subcommands: []cli.Command{
		{
			Name:  "up",
			Usage: "apply up migration direction",
			Action: func(c *cli.Context) error {
				config.Init(c.GlobalString(FlagConfig))

				err := migrateDB(db.MigrateUp)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "down",
			Usage: "drop and clean database schema",
			Action: func(c *cli.Context) error {
				config.Init(c.GlobalString(FlagConfig))

				err := migrateDB(db.MigrateDown)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "redo",
			Usage: "reset database schema",
			Action: func(c *cli.Context) error {
				config.Init(c.GlobalString(FlagConfig))

				err := migrateDB(db.MigrateDown)
				if err != nil {
					return err
				}

				err = migrateDB(db.MigrateUp)
				if err != nil {
					return err
				}

				return nil
			},
		},
	},
}

func migrateDB(direction db.MigrateDir) *cli.ExitError {
	count, err := dbschema.Migrate(config.Config().DB, direction)
	if err != nil {
		log.Default.WithError(err).Error("Migrations failed")
		return cli.NewExitError(fmt.Sprintf("migration %s failed", direction), 1)
	}

	log.Default.Info(fmt.Sprintf("Applied %d %s migration", count, direction))
	return nil
}
