package initialization

import (
	"fmt"
	"sync"
	"time"

	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/armory/tools"
	"github.com/lancer-kit/domain-based-scaffold/actions/migrations"
	"github.com/lancer-kit/domain-based-scaffold/config"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const flagConfig = "config"
const defaultInitInterval = 5 * time.Second

var initConfigs = map[initModule]func(*config.Configuration, *logrus.Entry) error{
	DB: initDatabase,
	//NATS: initNATS,
}

func Init(c *cli.Context) *config.Configuration {
	config.Init(c.GlobalString(flagConfig))
	cfg := config.Config()

	wg := sync.WaitGroup{}
	for module, initializer := range initConfigs {
		var timeout time.Duration
		if module == DB {
			timeout = time.Duration(cfg.DB.InitTimeout) * time.Second
		} else {
			timeout = time.Duration(cfg.ModulesInitTimeout) * time.Second
		}

		wg.Add(1)

		go func(module initModule, initializer func(*config.Configuration, *logrus.Entry) error, timeout time.Duration) {
			defer wg.Done()
			ok := tools.RetryIncrementallyUntil(
				defaultInitInterval,
				timeout,

				func() bool {
					err := initializer(cfg, log.Default)
					if err != nil {
						log.Default.WithError(err).Error("Can't init " + module)
					}
					return err == nil
				})
			if !ok {
				log.Default.Fatal("Can't init " + module)
			}
		}(module, initializer, timeout)
	}

	wg.Wait()

	if cfg.DB.AutoMigrate {
		count, err := migrations.Migrate(config.Config().DB.ConnURL, "up")
		if err != nil {
			log.Default.WithError(err).Fatal("Migrations failed")
			return cfg
		}

		log.Default.Info(fmt.Sprintf("Applied %d %s migration", count, "up"))
	}

	return cfg
}
