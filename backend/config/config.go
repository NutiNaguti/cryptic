package config

import (
	"os"
	"strconv"
)

type Config struct {
	RPC_URL       string
	ERC20_ADDRESS string
	DEBUG_MODE    bool
}

func New() *Config {
	return &Config{
		RPC_URL:       getEnv("RPC_URL", ""),
		ERC20_ADDRESS: getEnv("ERC20_ADDRESS", ""),
		DEBUG_MODE:    getEnvAsBool("DEBUG_MODE", true),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsBool(key string, defaultVal bool) bool {
	valStr := getEnv(key, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
