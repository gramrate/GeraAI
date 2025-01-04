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
			env.GetEnv("POSTGRES_HOST", ""),
			env.GetEnv("POSTGRES_USER", ""),
			env.GetEnv("POSTGRES_PASSWORD", ""),
			env.GetEnv("POSTGRES_DB", ""),
			env.GetEnv("POSTGRES_PORT", "")),
		JWTSecret:     env.GetEnv("JWT_SECRET", ""),
		JWTExpiration: time.Hour * 24 * 30,
		ApiKey:        env.GetEnv("OPENAI_API_KEY", ""),
		ProxyURL:      env.GetEnv("PROXY_URL", ""),
	}
	fmt.Println(Config.DBConnectionString)
}
