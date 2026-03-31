package tests

import (
	"desafio-cloud-run/handlers"
	"desafio-cloud-run/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIsValidCEP(t *testing.T) {
	testCases := []struct {
		cep      string
		expected bool
	}{
		{"12345678", true},
		{"1234567", false},
		{"123456789", false},
		{"abcdefgh", false},
		{"1234567a", false},
	}

	for _, tc := range testCases {
		result := utils.IsValidCEP(tc.cep)
		if result != tc.expected {
			t.Errorf("For CEP '%s', expected %v but got %v", tc.cep, tc.expected, result)
		}
	}
}

type MockLocationProvider struct{}

func (m *MockLocationProvider) GetLocation(cep string) (string, error) {
	if cep == "50070095" {
		return "Recife", nil
	}
	return "", nil
}

type MockWeatherProvider struct{}

func (m *MockWeatherProvider) GetWeather(city string) (float64, error) {
	if city == "Recife" {
		return 28.0, nil
	}
	return 0, nil
}

func TestWeatherHandler(t *testing.T) {
	locationProvider := &MockLocationProvider{}
	weatherProvider := &MockWeatherProvider{}

	handler := &handlers.WeatherHandler{
		LocationProvider: locationProvider,
		WeatherProvider:  weatherProvider,
	}

	req, err := http.NewRequest("GET", "/weather?cep=50070095", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"temp_C":28,"temp_F":82.4,"temp_K":301}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
