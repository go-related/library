package configuration

import (
	"errors"
	log "github.com/jt/books/gen/log"
	"github.com/spf13/viper"
)

type LibraryConfiguration struct {
	DbConnectionString string
}

var ApplicationConfiguration LibraryConfiguration

func LoadConfiguration(logger *log.Logger) {
	v := viper.New()
	v.SetConfigName("default")
	v.SetConfigType("yml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		// It's okay if there isn't a config file
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.Is(err, &configFileNotFoundError) {
			logger.Error().Err(err).Msg("error loading config file")
		}
	}
	v.AutomaticEnv()
	err := v.UnmarshalExact(&ApplicationConfiguration)
	if err != nil {
		logger.Error().Err(err).Msg("error mapping config file to type")
	}
}
