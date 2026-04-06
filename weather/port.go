package weather

import (
	"weather-api/domain"
)

type Service interface {
	GetWeather(city string) (*domain.WeatherResponse, error)
}

type WeatherClient interface {
	GetWeather(city string) (*domain.WeatherResponse, error)
}

type CacheClient interface {
	Get(city string) (*domain.WeatherResponse, error)
	Set(city string, weather *domain.WeatherResponse) error
}
