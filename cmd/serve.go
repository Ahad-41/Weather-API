package cmd

import (
	"weather-api/config"
	"weather-api/external/weatherapi"
	"weather-api/infra/redis"
	"weather-api/rest"
	weatherHandler "weather-api/rest/handlers/weather"
	"weather-api/weather"
)

func Serve() {
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
