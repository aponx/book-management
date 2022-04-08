package common

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Server struct {
	Host string
	Port string
}

type Log struct {
	Level  string
	Pretty bool
}

type DB struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

type Config struct {
	Server Server
	Log    Log
	DB     DB
}

var GlobalConfig *Config

func NewConfig() (conf *Config, err error) {
	defaultConfig := Config{}
	defaultConfigJson, err := json.Marshal(&defaultConfig)

	if err != nil {
		return
	}

	v := viper.NewWithOptions(viper.KeyDelimiter("_"))
	v.AddConfigPath(".")
	v.SetConfigFile(".env")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = v.ReadInConfig()
	if err != nil {
		log.Err(err).Msgf("Config file is not found")
		v.SetConfigType("json")

		err = v.ReadConfig(bytes.NewReader(defaultConfigJson))
		if err != nil {
			return
		}
		v.AllowEmptyEnv(true)
		v.AutomaticEnv()
	} else {
		log.Info().Msgf("using config file %s", v.ConfigFileUsed())
	}

	err = v.Unmarshal(&conf)
	if err != nil {
		return
	}

	return
}

func (l *Log) IsDebug() bool {
	return strings.ToLower(l.Level) == "debug"
}
