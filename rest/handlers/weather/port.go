package weather

import (
	"weather-api/domain"
)

type Service interface {
	GetWeather(city string) (*domain.WeatherResponse, error)
}
