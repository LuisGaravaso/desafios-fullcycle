package cepparser

import (
	"fmt"
	"strings"
)

func ParseCep(cep string) (string, error) {
	cep = strings.ReplaceAll(cep, "-", "")
	if len(cep) != 8 || strings.ContainsAny(cep, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return "", fmt.Errorf(cep)
	}
	return cep, nil
}
