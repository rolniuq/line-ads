package oauth2

import (
	"line-ads/configs"
	"line-ads/pkgs/logger"

	"github.com/submodule-org/submodule.go/v2"
)

var Oauth2Mod = submodule.Make[Oauth2](func(config *configs.Config, logger *logger.Logger) Oauth2 {
	return NewOauth2(config, logger)
}, configs.ConfigMod, logger.LoggerMod)

type Oauth2 interface {
	GetPresignedUrl() (string, error)
}

type oauth2 struct {
	config *configs.Config
	logger *logger.Logger
}

func NewOauth2(config *configs.Config, logger *logger.Logger) Oauth2 {
	return &oauth2{config, logger}
}

func (o *oauth2) GetPresignedUrl() (string, error) {
	return "", nil
}
