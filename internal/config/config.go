package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Env      string
	GrpcPort int
	MongoDns string
	HttpPort int
}

func Load() (*Config, error) {
	var cfg Config
	if err := godotenv.Load(); err != nil {
		return &cfg, err
	}

	grpcPort, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		return &cfg, err
	}

	httpPort, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return &cfg, err
	}

	cfg.MongoDns = os.Getenv("MONGO_DNS")
	cfg.GrpcPort = grpcPort
	cfg.HttpPort = httpPort

	return &cfg, nil
}
