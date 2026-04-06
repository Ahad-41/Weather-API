package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"weather-api/domain"
	"weather-api/weather"
)

type redisCache struct {
	client *RedisDB
}

func NewCacheClient(client *RedisDB) weather.CacheClient {
	if client == nil {
		return nil
	}
	return &redisCache{client: client}
}

func (r *redisCache) Get(city string) (*domain.WeatherResponse, error) {
	if r.client == nil {
		return nil, nil
	}
	key := fmt.Sprintf("weather:%s", city)
	val, err := r.client.Client.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	var weatherResponse domain.WeatherResponse
	err = json.Unmarshal([]byte(val), &weatherResponse)
	if err != nil {
		return nil, err
	}

	return &weatherResponse, nil
}

func (r *redisCache) Set(city string, weather *domain.WeatherResponse) error {
	if r.client == nil {
		return nil
	}
	key := fmt.Sprintf("weather:%s", city)
	val, err := json.Marshal(weather)
	if err != nil {
		return err
	}

	// 12 hours expiration
	err = r.client.Client.Set(context.Background(), key, val, 12*time.Hour).Err()
	return err
}
