package handlers

import (
	"desafio-cloud-run/interfaces"
	"desafio-cloud-run/models"
	"desafio-cloud-run/utils"
	"encoding/json"
	"net/http"
)

type WeatherHandler struct {
	LocationProvider interfaces.LocationProvider
	WeatherProvider  interfaces.WeatherProvider
}

func (h *WeatherHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if !utils.IsValidCEP(cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	city, err := h.LocationProvider.GetLocation(cep)
	if err != nil {
		if err.Error() == "can not find zipcode" {
			http.Error(w, "can not find zipcode", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	tempC, err := h.WeatherProvider.GetWeather(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tempF := tempC*1.8 + 32
	tempK := tempC + 273

	response := models.WeatherResponse{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
