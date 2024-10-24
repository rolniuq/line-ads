package configs

import (
	"fmt"
	"line-ads/pkgs/path"

	"github.com/spf13/viper"
	"github.com/submodule-org/submodule.go/v2"
)

type Config struct {
	ClientId     string
	ClientSecret string
}

var ConfigMod = submodule.Make[*Config](func() (*Config, error) {
	root := path.RootDir()
	viper.SetConfigFile(fmt.Sprintf("%s/config.yaml", root))
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		ClientId:     viper.GetString("line.client_key"),
		ClientSecret: viper.GetString("line.client_secret"),
	}, nil
})
