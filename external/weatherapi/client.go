package weatherapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
	"weather-api/domain"
)

type WeatherClient struct {
	BaseURL string
	APIKey  string
	HTTPClient *http.Client
}

func NewWeatherClient(apiKey string) *WeatherClient {
	return &WeatherClient{
		BaseURL: "http://api.weatherapi.com/v1/current.json",
		APIKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *WeatherClient) GetWeather(city string) (*domain.WeatherResponse, error) {
	escapedCity := url.QueryEscape(city)
	fullURL := fmt.Sprintf("%s?key=%s&q=%s&aqi=no", c.BaseURL, c.APIKey, escapedCity)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	var weatherResponse domain.WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &weatherResponse, nil
}
