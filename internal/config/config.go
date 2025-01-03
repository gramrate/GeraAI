package config

import (
	"fmt"
	"gera-ai/internal/utils/dbURL"
	"gera-ai/internal/utils/env"
	"time"
)

var (
	Config AppConfig = AppConfig{}
)

type AppConfig struct {
	DBConnectionString string
	JWTSecret          string
	JWTExpiration      time.Duration
	ApiKey             string
	ProxyURL           string
}

func InitConfig() {
	Config = AppConfig{
		DBConnectionString: dbURL.GetDbUrl(
			env.GetEnv("DB_HOST", "db"),
			env.GetEnv("DB_USER", "postgres"),
			env.GetEnv("DB_PASSWORD", "postgres"),
			env.GetEnv("DB_NAME", "geraai"),
			env.GetEnv("DB_PORT", "5432")),
		JWTSecret:     env.GetEnv("JWT_SECRET", ""),
		JWTExpiration: time.Hour * 24 * 30,
		ApiKey:        env.GetEnv("OPENAI_API_KEY", ""),
		ProxyURL:      env.GetEnv("PROXY_URL", ""),
	}
	fmt.Println(Config.DBConnectionString)
}
