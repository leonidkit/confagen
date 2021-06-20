package main

import (
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	App   App
	Other Other
}

func New(file string) (*Config, error) {
	vip := viper.New()
	vip.SetConfigFile(file)

	if err := vip.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "unable to read config file")
	}

	return &Config{
		App: App{
			Port:         vip.GetInt("app.port"),
			LogLevel:     vip.GetString("app.log_level"),
			ReadTimeout:  vip.GetDuration("app.read_timeout"),
			WriteTimeout: vip.GetDuration("app.write_timeout"),
		},
		Other: Other{
			Tick:    vip.GetDuration("other.tick"),
			ApiUrl:  vip.GetString("other.api.url"),
			ApiName: vip.GetString("other.api.name"),
			Url:     vip.GetString("other.url"),
		},
	}, nil
}

type App struct {
	Port         int
	LogLevel     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Other struct {
	Tick    time.Duration
	ApiUrl  string
	ApiName string
	Url     string
}
