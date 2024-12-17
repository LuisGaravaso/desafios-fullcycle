package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"findceps/models"
	"findceps/services"
)

type ChannelMessage struct {
	Message *http.Response
	Service string
}

func ProcessCep(cep, outputType string, wg *sync.WaitGroup) {
	defer wg.Done()

	results := make(chan ChannelMessage, 2)
	errors := make(chan error, 2)

	go func() {
		data, err := services.MakeRequest(cep, "brasilapi")
		if err != nil {
			errors <- err
		} else {
			results <- ChannelMessage{Message: data, Service: "brasilapi"}
		}
	}()

	go func() {
		data, err := services.MakeRequest(cep, "viacep")
		if err != nil {
			errors <- err
		} else {
			results <- ChannelMessage{Message: data, Service: "viacep"}
		}
	}()

	select {
	case r := <-results:
		defer r.Message.Body.Close()

		if outputType == "-raw" {
			if _, err := io.Copy(os.Stdout, r.Message.Body); err != nil {
				fmt.Println("Erro ao copiar a resposta:", err)
			}
		} else {
			var message models.MessageInterface
			switch r.Service {
			case "viacep":
				var data models.ViacepMessage
				if err := json.NewDecoder(r.Message.Body).Decode(&data); err != nil {
					fmt.Printf("Erro ao decodificar JSON do %s: %s", r.Service, err)
					return
				}
				data.Cep = cep
				message = data
			case "brasilapi":
				var data models.BrasilAPIMessage
				if err := json.NewDecoder(r.Message.Body).Decode(&data); err != nil {
					fmt.Printf("Erro ao decodificar JSON do %s: %s", r.Service, err)
					return
				}
				data.Cep = cep
				message = data
			}
			cepMessage := message.Translate()
			if outputType == "-text" {
				fmt.Println(cepMessage.String())
			} else {
				json.NewEncoder(os.Stdout).Encode(cepMessage)
			}
		}
	case err := <-errors:
		fmt.Printf("Erro ao buscar o CEP %s: %v\n", cep, err)
	case <-time.After(1 * time.Second):
		fmt.Printf("Erro: Timeout para o CEP %s - Nenhuma API respondeu a tempo\n", cep)
	}
}
