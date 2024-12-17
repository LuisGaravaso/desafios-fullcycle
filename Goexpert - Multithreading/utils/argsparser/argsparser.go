package argparser

import (
	"encoding/json"
	"findceps/models"
	"findceps/utils/cepparser"
	"fmt"
	"os"
)

func ParseArgs(args []string) (ceps []string, invalidCeps []string, outputType string) {
	outputType = "-json"
	nArgs := len(args)
	ceps = make([]string, nArgs-1)
	validIndex := 0

	invalidCeps = make([]string, 0)

	for _, arg := range args[1:] {
		if arg == "-text" || arg == "-raw" {
			outputType = arg
		} else {
			parsedCep, err := cepparser.ParseCep(arg)
			if err != nil {
				invalidCeps = append(invalidCeps, err.Error())
				continue
			}
			ceps[validIndex] = parsedCep
			validIndex++
		}
	}
	ceps = ceps[:validIndex]

	return
}

func OutputInvalidCeps(invalidCeps []string, outputType string) {
	for _, invalidCep := range invalidCeps {
		CepMessage := models.CepMessage{Cep: invalidCep, Service: "invalid"}
		switch outputType {
		case "-text":
			fmt.Println(CepMessage.String())
		default:
			json.NewEncoder(os.Stdout).Encode(CepMessage)
		}
	}
}
