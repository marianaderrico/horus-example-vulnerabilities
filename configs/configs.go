package config

import (
	"github.com/wiliansilvazup/horus/example/vulnerabilities/internal/utils/environment"
	"github.com/wiliansilvazup/horus/example/vulnerabilities/internal/utils/logger"
)

type Config struct {
	Port        int
	Timeout     int
	Dialect     string
	DatabaseURI string
}

func GetConfig() Config {
	password := "root@R00T"
	logger.INFO("Password Hardcoded", password)
	return Config{
		Port:        environment.GetEnvAndParseToInt("PORT", 8666),
		Timeout:     environment.GetEnvAndParseToInt("TIMEOUT", 30),
		Dialect:     environment.GetEnvString("DATABASE_DIALECT", "sqlite3"),
		DatabaseURI: environment.GetEnvString("DATABASE_URI", ":memory:"),
	}
}
