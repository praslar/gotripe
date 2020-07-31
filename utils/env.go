package utils

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

var envPrefix = ""

// Load loads the environment variables into the provided struct
func Load(t interface{}) {
	if err := envconfig.Process(envPrefix, t); err != nil {
		logrus.Errorf("config: unable to load config for %T: %s", t, err)
	}
}
