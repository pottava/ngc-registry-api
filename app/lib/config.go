package lib

import (
	"github.com/kelseyhightower/envconfig"
)

// Config can be set via environment variables
type config struct {
	APIEndpoint string `envconfig:"API_ENDPOINT" default:"http://localhost:9000"`
	AccessLog   bool   `envconfig:"ACCESS_LOG" default:"true"`
}

// Config represents its configurations
var Config *config

func init() {
	cfg := &config{}
	envconfig.MustProcess("ngc", cfg)
	Config = cfg
}
