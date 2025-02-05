package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	RedisAddr string `envconfig:"REDIS_ADDR" default:"redis:6379"`
	KafkaAddr string `envconfig:"KAFKA_ADDR" default:"kafka:9092"`
}

func New() *Config {
	return &Config{}
}

func (r *Config) Init() {
	if err := envconfig.Process("", r); err != nil {
		log.Fatalf("failed to load configuration: %s", err)
	}
}
