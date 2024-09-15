package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Creds struct {
	SERVER_PORT int `env:"SERVER_PORT"`
	DB_HOST     string `env:DB_HOST`
	DB_USER     string `env:DB_USER`
	DB_PASSWORD string `env:DB_PASSWORD`
	DB_NAME     string `env:DB_NAME`
	DB_PORT     int    `env:DB_PORT`
	DB_SSL_MODE string `env:DB_SSL_MODE`
	DB_TIMEZONE string `env:DB_TIMEZONE`
}

// load config through viper
func LoadConfig() *Creds {
	log.Info("Loading Configs......")
	creds := Creds{}

	//set the config file
	viper.SetConfigFile(".env")

	//find and read the file defined in config path
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Error Reading Config File...Exiting..")
	}

	err = viper.Unmarshal(&creds)
	if err != nil {
		log.Error("Error in Unmarshalling config...Exiting..")
	}
	log.Info("Loading Configs : Success......")
	return &creds
}
