package interfaces

type WeatherProvider interface {
	GetWeather(city string) (float64, error)
}
