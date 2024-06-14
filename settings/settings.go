package settings

import (
	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	ArgoCDNamespace         string `envconfig:"ARGOCD_NAMESPACE" required:"true" default:"argocd"`
	ArgoCDManagerSecretName string `envconfig:"ARGOCD_MANAGER_SECRET_NAME" required:"true" default:"argocd-manager"`
	DisableFeatures         bool   `envconfig:"DISABLE_FEATURES" required:"true" default:"false"`
	HTTPPort                string `envconfig:"HTTP_PORT" required:"true" default:"8080"`
}

func (s Settings) Validation() error {
	return nil
}

func NewSettings() (Settings, error) {
	var settings Settings

	err := envconfig.Process("", &settings)
	if err != nil {
		return settings, err
	}

	if settings.Validation() != nil {
		return settings, err
	}

	return settings, nil
}
