package config

import (
	"time"

	"github.com/spf13/viper"
)

type Configuration struct {
	AppEnv string `mapstructure:"APPLICATION_ENV"`
	AppPort string `mapstructure:"APPLICATION_PORT"`
	AppDefaultTimeout time.Duration `mapstructure:"APPLICATION_DEFAULT_TIMEOUT"`

	CandaanEndpoint     string `mapstructure:"CANDAAN_API_ENDPOINT"`
	JokesBapackEndpoint string `mapstructure:"JOKES_BAPACK_ENDPOINT"`
}

func LoadConfig(path string) *Configuration {
	vp := viper.New()
	vp.AddConfigPath(path)
	vp.AddConfigPath(".")
	vp.SetConfigName("app")
	vp.SetConfigType("env")
	vp.AutomaticEnv()

	if err := vp.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Configuration
	if err := vp.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
}