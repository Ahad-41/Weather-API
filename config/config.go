package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version                  string
	ServiceName              string
	HttpPort                 int
	ApiKey                   string
	RedisAddr                string
	RedisPassword            string
	RateLimitRequests        int
	RateLimitDurationSeconds int
}

var configurations *Config

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: Failed to load .env file, using existing environment variables")
	}

	version := getEnv("VERSION", "1.0.0")
	serviceName := getEnv("SERVICE_NAME", "weather-api")
	httpPortStr := getEnv("HTTP_PORT", "8080")
	apiKey := getEnv("API_KEY", "")
	redisAddr := getEnv("REDIS_ADDR", "localhost:6379")
	redisPassword := getEnv("REDIS_PASSWORD", "")
	rateLimitRequestsStr := getEnv("RATE_LIMIT_REQUESTS", "10")
	rateLimitDurationSecondsStr := getEnv("RATE_LIMIT_DURATION_SECONDS", "60")

	if apiKey == "" {
		fmt.Println("API_KEY is required")
		os.Exit(1)
	}

	httpPort, err := strconv.Atoi(httpPortStr)
	if err != nil {
		fmt.Println("HTTP_PORT must be a number")
		os.Exit(1)
	}

	rateLimitRequests, err := strconv.Atoi(rateLimitRequestsStr)
	if err != nil {
		fmt.Println("RATE_LIMIT_REQUESTS must be a number")
		os.Exit(1)
	}

	rateLimitDurationSeconds, err := strconv.Atoi(rateLimitDurationSecondsStr)
	if err != nil {
		fmt.Println("RATE_LIMIT_DURATION_SECONDS must be a number")
		os.Exit(1)
	}

	configurations = &Config{
		Version:                  version,
		ServiceName:              serviceName,
		HttpPort:                 httpPort,
		ApiKey:                   apiKey,
		RedisAddr:                redisAddr,
		RedisPassword:            redisPassword,
		RateLimitRequests:        rateLimitRequests,
		RateLimitDurationSeconds: rateLimitDurationSeconds,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
