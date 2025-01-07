package envs

import (
	"log"
	"sync"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Environment       string        `env:"ENVIRONMENT"`
	HTTPServerPort    string        `env:"HTTP_SERVER_PORT"`
	CorsMaxAge        int           `env:"CORS_MAX_AGE"`
	HTTPServerTimeout time.Duration `env:"HTTP_SERVER_TIMEOUT"`

	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresDB       string `env:"POSTGRES_DB"`
}

var (
	once      sync.Once
	EnvConfig Config
)

func Envs() {
	once.Do(func() {
		err := env.Parse(&EnvConfig)
		if err != nil {
			log.Fatalf("fail to load environment variables: %s", err.Error())
		}
		if EnvConfig == (Config{}) {
			log.Fatal("fail to load environment variables")
		}
	})
}
