package configuration

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type LibraryConfiguration struct {
	DbConnectionString string
}

var ApplicationConfiguration LibraryConfiguration

func init() {
	v := viper.New()
	v.SetConfigName("default")
	v.SetConfigType("yml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		// It's okay if there isn't a config file
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.Is(err, &configFileNotFoundError) {
			logrus.WithError(err).Errorf("error loading config file")
		}
	}
	v.AutomaticEnv()
	err := v.UnmarshalExact(&ApplicationConfiguration)
	if err != nil {
		logrus.WithError(err).Errorf("error mapping config file to type")
	}
}
