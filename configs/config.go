package configs

import (
	"github.com/spf13/viper"
	"github.com/submodule-org/submodule.go/v2"
)

type Config struct {
	ClientId     string
	ClientSecret string
}

var ConfigMod = submodule.Make[*Config](func() (*Config, error) {
	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		ClientId:     viper.GetString("line.client_key"),
		ClientSecret: viper.GetString("line.client_secret"),
	}, nil
})
