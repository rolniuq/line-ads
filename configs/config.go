package configs

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
	"github.com/submodule-org/submodule.go/v2"
)

type Config struct {
	ClientId     string
	ClientSecret string
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))

	return filepath.Dir(d)
}

var ConfigMod = submodule.Make[*Config](func() (*Config, error) {
	root := rootDir()
	viper.SetConfigFile(fmt.Sprintf("%s/config.yaml", root))
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		ClientId:     viper.GetString("line.client_key"),
		ClientSecret: viper.GetString("line.client_secret"),
	}, nil
})
