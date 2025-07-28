package auth

import (
	"line-ads/configs"
	"line-ads/pkgs/auth/oauth2"
	"line-ads/pkgs/logger"

	"github.com/submodule-org/submodule.go/v2"
)

var AuthMod = submodule.Make[Auth](func(client oauth2.Oauth2, config *configs.Config, logger *logger.Logger) Auth {
	return NewAuth(client, config, logger)
}, oauth2.Oauth2Mod, configs.ConfigMod, logger.LoggerMod)

type AuthType string

const (
	Oauth2 string = "oauth2"
)

type Auth interface {
	GetPresignedUrl() (string, error)
}

type auth struct {
	client Auth
	config *configs.Config
	logger *logger.Logger
}

func NewAuth(client oauth2.Oauth2, config *configs.Config, logger *logger.Logger) Auth {
	return &auth{client, config, logger}
}

func (a *auth) GetPresignedUrl() (string, error) {
	return "", nil
}
