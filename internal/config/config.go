package config

import (
	"github.com/elem1092/crud/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Configuration struct {
	Listen ServerConfig `yaml:"listen"`
	DBCfg  DBConfig     `yaml:"db"`
}

type ServerConfig struct {
	Port    string `yaml:"port"`
	Address string `yaml:"address"`
}

type DBConfig struct {
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Address     string `yaml:"address"`
	Port        string `yaml:"port"`
	DBName      string `yaml:"name"`
	MaxAttempts int    `yaml:"max_attempts"`
}

var cfg *Configuration = &Configuration{ServerConfig{}, DBConfig{}}
var once sync.Once

func GetConfiguration() *Configuration {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("Parsing configuration file")
		if err := cleanenv.ReadConfig("config.yml", cfg); err != nil {
			help, _ := cleanenv.GetDescription(cfg, nil)
			logger.Fatal(help)
			panic(err)
		}
		logger.Info("finished reading configuration file")
	})

	return cfg
}
