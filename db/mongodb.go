package db

import (
	"errors"
	"sync"
	"time"

	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"

	"github.com/gotripe/utils"
)

type (
	// Config hold MongoDB configuration information
	Config struct {
		Addrs    []string      `envconfig:"MONGODB_ADDRS" default:"localhost:27017"`
		Database string        `envconfig:"MONGODB_DATABASE" default:"gotriple"`
		Username string        `envconfig:"MONGODB_USERNAME"`
		Password string        `envconfig:"MONGODB_PASSWORD"`
		Timeout  time.Duration `envconfig:"MONGODB_TIMEOUT" default:"10s"`
	}
)

var (
	session     *mgo.Session
	sessionOnce sync.Once
)

// LoadConfigFromEnv load mongodb configurations from environments
func LoadConfigFromEnv() *Config {
	var conf Config
	utils.Load(&conf)
	return &conf
}

// Dial dial to target server with Monotonic mode
func Dial(conf *Config) (*mgo.Session, error) {
	logrus.Info(conf.Addrs)
	logrus.Infof("Dialing to target MongoDB at: %v, Database: %v", conf.Addrs, conf.Database)
	ms, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    conf.Addrs,
		Database: conf.Database,
		Username: conf.Username,
		Password: conf.Password,
		Timeout:  conf.Timeout,
	})
	if err != nil {
		return nil, err
	}
	ms.SetMode(mgo.Monotonic, true)
	logrus.Infof("successfully dialing to MongoDB at %v", conf.Addrs)
	return ms, nil
}

func IsErrNotFound(err error) bool {
	return errors.Is(err, mgo.ErrNotFound)
}

func DialDefaultMongoDB() (*mgo.Session, error) {
	repoConf := LoadConfigFromEnv()
	var err error
	sessionOnce.Do(func() {
		session, err = Dial(repoConf)
	})
	if err != nil {
		return nil, err
	}
	s := session.Clone()
	return s, nil
}
