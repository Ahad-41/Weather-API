package cmd

import (
	"fmt"
	"os/exec"
	"weather-api/config"
	"weather-api/external/weatherapi"
	"weather-api/infra/redis"
	"weather-api/rest"
	weatherHandler "weather-api/rest/handlers/weather"
	"weather-api/weather"
)

func Serve() {
	// ⚡ Automate Redis Startup
	fmt.Println("🚀 Ensuring Redis is running in Docker...")
	cmd := exec.Command("/bin/bash", "./docker-run-redis.sh")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("⚠️ Error starting Redis script: %v\nOutput: %s\n", err, string(output))
	} else {
		fmt.Println("✅ Redis startup script executed successfully.")
	}

	cnf := config.GetConfig()

	// 1. Initialize Infrastucture (Redis)
	redisClient := redis.ConnectRedis(cnf.RedisAddr, cnf.RedisPassword)
	cacheClient := redis.NewCacheClient(redisClient)

	// 2. Initialize External Clients (Weather API)
	weatherClient := weatherapi.NewWeatherClient(cnf.ApiKey)

	// 3. Initialize Business Logic (Weather Service)
	weatherService := weather.NewService(weatherClient, cacheClient)

	// 4. Initialize Presentatation layer (Handlers)
	handler := weatherHandler.NewHandler(weatherService)

	// 5. Initialize & Start Server
	server := rest.NewServer(cnf, handler)
	server.Start()
}
