package models

import (
	"findceps/utils/cepparser"
	"fmt"
)

type ViacepMessage struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Unidade     string `json:"unidade"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	DDD         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type BrasilAPIMessage struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

type CepMessage struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Estado     string `json:"estado"`
	Cidade     string `json:"cidade"`
	Bairro     string `json:"bairro"`
	Service    string `json:"service"`
}

func (c CepMessage) String() string {
	if c.Service == "invalid" {
		return fmt.Sprintf("CEP %s no Formato Errado \nNão foi consultado \n------------", c.Cep)
	}

	if c.Logradouro == "" {
		return fmt.Sprintf("CEP %s Inválido \nDados Obtidos de %s \n------------", c.Cep, c.Service)
	}

	return fmt.Sprintf(
		"%s, Bairro %s \n%s-%s \n%s \nDados Obtidos de %s \n------------",
		c.Logradouro, c.Bairro, c.Cidade, c.Estado, c.Cep, c.Service,
	)
}

func (v ViacepMessage) Translate() CepMessage {

	return CepMessage{
		Cep:        v.Cep,
		Logradouro: v.Logradouro,
		Estado:     v.Uf,
		Cidade:     v.Localidade,
		Bairro:     v.Bairro,
		Service:    "viacep",
	}
}

func (b BrasilAPIMessage) Translate() CepMessage {
	parsedCep, _ := cepparser.ParseCep(b.Cep)
	b.Cep = parsedCep[:5] + "-" + parsedCep[5:] // Adiciona o hífen no CEP

	return CepMessage{
		Cep:        b.Cep,
		Logradouro: b.Street,
		Estado:     b.State,
		Cidade:     b.City,
		Bairro:     b.Neighborhood,
		Service:    "brasilapi",
	}
}

type MessageInterface interface {
	Translate() CepMessage
}
