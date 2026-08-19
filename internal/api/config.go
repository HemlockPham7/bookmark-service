package api

import (
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppPort     string `default:"8080" envconfig:"APP_PORT"`
	LogLevel    string `default:"info" envconfig:"LOG_LEVEL"`
	BasePath    string `default:"/" envconfig:"BASE_PATH"`
	ServiceName string `default:"bookmark-service" envconfig:"SERVICE_NAME"`
	InstanceID  string `default:"" envconfig:"INSTANCE_ID"`
	QueueName   string `default:"bookmark-import" envconfig:"QUEUE_NAME"`
}

// NewConfig creates a new config
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("bookmark", cfg)
	if err != nil {
		return nil, err
	}

	if cfg.InstanceID == "" {
		cfg.InstanceID = uuid.New().String()
	}
	return cfg, err
}
