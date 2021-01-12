package config

import (
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Config appConfig

type appConfig struct {
	DB         *sqlx.DB
	DB_ERR     error
	RESTAPIPort        int    `mapstructure:"rest_api_port"`
	DBURL        string `mapstructure:"mysql_url"`
	LOGGER             *zap.SugaredLogger
}

func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("server")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("restful")

	v.AutomaticEnv()

	v.SetDefault("rest_api_port", 7080)

	v.SetDefault("mysql_url", "")



	return v.Unmarshal(&Config)
}
