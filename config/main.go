package config

import (
	"io/ioutil"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// config is a `Cfg` singleton var,
// for access use the `Config` method.
var config *Cfg

func Init(path string) {
	rawConfig, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.New().
			WithError(err).
			WithField("path", path).
			Fatal("unable to read config file")
	}

	config = new(Cfg)
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

	if config.Services == nil {
		config.FillDefaultServices()
	}

}

// Config returns the config obj.
func Config() *Cfg {
	return config
}
