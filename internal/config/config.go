package config

import (
    "github.com/elem1092/crud/pkg/logging"
    "github.com/ilyakaznacheev/cleanenv"
    "sync"
)

type Configuration struct {
    Listen struct {
        Port    string `yaml:"port"`
        Address string `yaml:"address"`
    } `yaml:"listen"`
    DB struct {
        Port     string `yaml:"port"`
        Address  string `yaml:"address"`
        Username string `yaml:"username"`
        Password string `yaml:"password"`
    } `yaml:"db"`
}

var cfg *Configuration
var once sync.Once

func GetConfiguration() *Configuration {
    once.Do(func() {
        logger := logging.GetLogger()
        logger.Info("Parsing configuration file")
        cfg := &Configuration{}
        if err := cleanenv.ReadConfig("config.yml", cfg); err != nil {
            help, _ := cleanenv.GetDescription(cfg, nil)
            logger.Fatal(help)
            panic(err)
        }
    })

    return cfg
}
