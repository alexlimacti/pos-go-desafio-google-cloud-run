package main

import (
	"desafio-cloud-run/handlers"
	"desafio-cloud-run/providers"
	"fmt"
	"net/http"
)

func main() {
	locationProvider := &providers.ViaCEPProvider{}
	weatherProvider := &providers.WeatherAPIProvider{}

	weatherHandler := &handlers.WeatherHandler{
		LocationProvider: locationProvider,
		WeatherProvider:  weatherProvider,
	}

	http.Handle("/weather", weatherHandler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
