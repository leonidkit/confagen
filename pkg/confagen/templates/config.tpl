package main

import (
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
{{ range .Struct }}	{{ .Name }} {{ .Name }}
{{ end }} }

func New(file string) (*Config, error) {
	vip := viper.New()
	vip.SetConfigFile(file)

	if err := vip.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "unable to read config file")
	}

	return &Config{
		{{ range .Struct }}	{{ .Name }}: {{ .Name }}{
                {{ range .Fields }}	{{ .Name }}: vip.Get{{ .ViperType }}("{{ .Path }}"),
				{{ end }} },
		{{ end }} }, nil
}

{{ range .Struct }}
type {{ .Name }} struct {
    {{ range .Fields }}{{ .Name }} {{ if eq .ValType "Duration" }} time.{{ .ValType }} {{ else }} {{ .ValType }} {{ end }}
    {{ end }} }
{{ end }}