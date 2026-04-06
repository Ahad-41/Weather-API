package weather

import (
	"fmt"
	"weather-api/domain"
)

type weatherService struct {
	weatherClient WeatherClient
	cacheClient   CacheClient
}

func NewService(weatherClient WeatherClient, cacheClient CacheClient) Service {
	return &weatherService{
		weatherClient: weatherClient,
		cacheClient:   cacheClient,
	}
}

func (s *weatherService) GetWeather(city string) (*domain.WeatherResponse, error) {
	// 1. Try to get from the cache
	if s.cacheClient != nil {
		cachedWeather, err := s.cacheClient.Get(city)
		if err == nil && cachedWeather != nil {
			fmt.Printf("Cache hit for city: %s\n", city)
			return cachedWeather, nil
		}
		if err != nil {
			fmt.Printf("Cache miss or error for city %s: %v\n", city, err)
		}
	}

	// 2. Not in cache, fetch from the 3rd party API
	weather, err := s.weatherClient.GetWeather(city)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather for %s from API: %w", city, err)
	}

	// 3. Save to cache for next time
	if s.cacheClient != nil {
		err = s.cacheClient.Set(city, weather)
		if err != nil {
			fmt.Printf("Warning: Failed to save weather to cache for city %s: %v\n", city, err)
		} else {
			fmt.Printf("Weather saved to cache for city: %s\n", city)
		}
	}

	return weather, nil
}
