package config

import (
	"io/ioutil"

	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/uwe/v2"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	// TODO: change me
	ServiceName = "domain-based-scaffold"

	WorkerAPIServer uwe.WorkerName = "api-server"
	WorkerFooBar    uwe.WorkerName = "foobar"
)

var availableWorkers = map[uwe.WorkerName]struct{}{
	WorkerAPIServer: {},
	WorkerFooBar:    {},
}

// config is a `Configuration` singleton var,
// for access use the `Config` method.
var config *Configuration

func Init(path string) {
	rawConfig, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.New().
			WithError(err).
			WithField("path", path).
			Fatal("unable to read config file")
	}

	config = new(Configuration)
	err = yaml.Unmarshal(rawConfig, config)
	if err != nil {
		logrus.New().
			WithError(err).
			WithField("raw_config", rawConfig).
			Fatal("unable to unmarshal config file")
	}

	err = config.Validate()
	if err != nil {
		logrus.New().
			WithError(err).
			Fatal("Invalid configuration")
	}

	initLog()
}

// Config returns the config obj.
func Config() *Configuration {
	return config
}

func initLog() {
	_, err := log.Init(config.Log)
	if err != nil {
		log.Default.
			WithError(err).
			Fatal("Unable to init log")
	}
}
