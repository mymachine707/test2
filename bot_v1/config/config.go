package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	App         string
	AppVersion  string
	Environment string // development, staging, orderion
	HTTPPort    string
	GRPCPort    string

	AdminServiceGrpcHost string
	AdminServiceGrpcPort string
}

// Load ...
func Load() Config {
	// .env file bor yo'qligini tekshirilvotti
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":7070"))

	config.App = cast.ToString(getOrReturnDefaultValue("APP", "article"))
	config.AppVersion = cast.ToString(getOrReturnDefaultValue("APP_VERSION", "1.0.0"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "development"))

	config.GRPCPort = cast.ToString(getOrReturnDefaultValue("GRPC_PORT", ":9001"))

	config.AdminServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("CATEGORY_SERVICE_HOST", "localhost"))
	config.AdminServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("CATEGORY_SERVICE_PORT", ":9090"))
	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
