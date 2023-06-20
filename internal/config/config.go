package config

import (
	"sync"

	"github.com/xxlifestyle/notes_auth-service/pkg/logging"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
}

var instance *Config

var once sync.Once

func GetConfig() *Config {
	once.Do(
		func() {
			logger := logging.GetLogger()
			logger.Info("reading config")
			instance = &Config{}
			err := cleanenv.ReadConfig("config.yml", instance)

			if err != nil {
				logger.Errorf("error during reading .yml config, error: %v", err)
			}
		})
	return instance
}
