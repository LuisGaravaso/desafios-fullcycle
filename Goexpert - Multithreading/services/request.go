package services

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	ViaCepURL      = "https://viacep.com.br/ws/"
	BrasilAPIURL   = "https://brasilapi.com.br/api/cep/v1/"
	RequestTimeout = 1 * time.Second
)

func MakeRequest(cep, service string) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RequestTimeout)
	defer cancel()

	var url string
	switch service {
	case "viacep":
		url = fmt.Sprintf("%s%s/json/", ViaCepURL, cep)
	case "brasilapi":
		url = fmt.Sprintf("%s%s", BrasilAPIURL, cep)
	default:
		return nil, fmt.Errorf("serviço não suportado: %s", service)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	return client.Do(req)
}
