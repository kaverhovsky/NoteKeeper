package common

import (
	"github.com/asaskevich/govalidator"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

// Config - struct that stores configuration of the app
type Config struct {
	Debug       bool   `envconfig:"DEBUG" default:"false" valid:"optional,type(bool)"`
	Development bool   `envconfig:"DEVELOPMENT" default:"false" valid:"optional,type(bool)"`
	LogLevel    string `envconfig:"LOGLEVEL" default:"info" valid:"in(debug|info|warning|error)"`
	Listen      string `envconfig:"LISTEN" valid:"required,dialstring"`
	Postgres    string `envconfig:"DATABASE_POSTGRES" valid:"required,requrl"`
}

// ReadConfig returns *Config with app configuration
func ReadConfig(path string, log *zap.Logger) (*Config, error) {
	logger := log.Named("config")
	if err := loadConfig(path); err != nil {
		return nil, err
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	if _, err := govalidator.ValidateStruct(cfg); err != nil {
		return nil, err
	}

	logger.With(zap.Any("config", cfg)).Info("Config loaded")
	return &cfg, nil
}

func loadConfig(path string) error {
	if path == "" {
		return nil
	}
	if err := godotenv.Load(path); err != nil {
		return err
	}
	return nil
}

// Mode returns application mode: "production" or "development".
func (c Config) Mode() string {
	if c.Development {
		return "development"
	}
	return "production"
}

// GetLogLevel returns logging level with respect to application mode.
func (c Config) Level() string {
	if c.Debug {
		return "debug"
	}
	if c.LogLevel == "warning" {
		return "warn"
	}
	return c.LogLevel
}
