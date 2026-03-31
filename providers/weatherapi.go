package providers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const weatherAPIKey = "c9a09322384b4221980120536241107"

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

type WeatherAPIProvider struct {
	Client *http.Client
}

func (p *WeatherAPIProvider) GetWeather(city string) (float64, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", weatherAPIKey, url.QueryEscape(city))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	client := p.Client
	if client == nil {
		client = http.DefaultClient
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var weatherAPIResponse WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherAPIResponse); err != nil {
		return 0, err
	}

	return weatherAPIResponse.Current.TempC, nil
}
