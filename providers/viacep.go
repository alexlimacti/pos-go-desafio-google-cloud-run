package providers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ViaCEPResponse struct {
	Localidade string `json:"localidade"`
	Erro       bool   `json:"erro"`
}

type ViaCEPProvider struct {
	Client *http.Client
}

func (p *ViaCEPProvider) GetLocation(cep string) (string, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	client := p.Client
	if client == nil {
		client = http.DefaultClient
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var viaCEPResponse ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCEPResponse); err != nil {
		return "", err
	}

	if viaCEPResponse.Erro {
		return "", fmt.Errorf("can not find zipcode")
	}

	return viaCEPResponse.Localidade, nil
}
