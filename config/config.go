package config

import (
	"log"

	env "github.com/Netflix/go-env"
)

var cfg Config

type Config struct {
	AppConfig      AppConfig
	PostgresConfig PostgresConfig
	// DbPostgresConfig DbPostgresConfig
	RedisConfig RedisConfig
}

type AppConfig struct {
	Port string `env:"APP_PORT,required=true"`
}

type PostgresConfig struct {
	PostgresHost     string `env:"POSTGRES_HOST,required=true"`
	PostgresUser     string `env:"POSTGRES_USER,required=true"`
	PostgresPassword string `env:"POSTGRES_PASSWORD,required=true"`
	PostgresName     string `env:"POSTGRES_NAME,required=true"`
	PostgresPort     string `env:"POSTGRES_PORT,required=true"`
	PostgresSSLMode  string `env:"POSTGRES_SSLMODE,required=true"`
	PostgresTimezone string `env:"POSTGRES_TIMEZONE,required=true"`
}

//	type DbPostgresConfig struct{
//		PostgresHost     string `env:"POSTGRES_HOST,required=true"`
//		PostgresUser     string `env:"POSTGRES_USER,required=true"`
//		PostgresPassword string `env:"POSTGRES_PASSWORD,required=true"`
//		PostgresName     string `env:"POSTGRES_NAME,required=true"`
//		PostgresPort     string `env:"POSTGRES_PORT,default=5432"`
//		PostgresSSLMode  string `env:"POSTGRES_SSLMODE,default=disable"`
//		PostgresTimezone string `env:"POSTGRES_TIMEZONE,default=Asia/Shanghai"`
//	}
type RedisConfig struct {
	RedisAddr     string `env:"REDIS_ADDR,required=true"`
	RedisPassword string `env:"REDIS_PASSWORD,required=true"`
	RedisDb       int `env:"REDIS_DB,required=true"`
}

func init() {
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Panic(err)
	}
}

func GetConfig() Config {
	return cfg
}
