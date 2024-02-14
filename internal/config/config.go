package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Env string `yaml:"env"`
	GrpcPort int `yaml:"grpc_port"`
	MongoUrl string `yaml:"mongo_url"`
	HttpAddress string `yaml:"http_address"`

}

func Load() *Config {
	data, err := os.ReadFile(getFilePath())
    if err != nil {
        panic(err)
    }

	var cfg Config
	
	yaml.Unmarshal(data, &cfg)
	return &cfg
}

func getFilePath() string {
	currentDir, err := os.Getwd()
    if err != nil {
        panic(err)
    }

	return filepath.Join(currentDir, "config/config.yaml")
}