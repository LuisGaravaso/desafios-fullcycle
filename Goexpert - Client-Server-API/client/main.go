package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Cotacao struct {
	Valor float64 `json:"valor"`
}

const urlCotacao = "http://localhost:8080/cotacao"

func main() {
	// Cria o contexto de requisição com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// Realiza a requisição para obter a cotação
	cotacao, err := obterCotacao(ctx)
	if err != nil {
		log.Fatalf("Erro ao obter cotação: %v", err)
	}

	// Escreve a cotação em um arquivo
	if err := salvarCotacao(cotacao); err != nil {
		log.Fatalf("Erro ao salvar cotação no arquivo: %v", err)
	}

	log.Println("Cotação salva com sucesso em 'cotacao.txt'")
}

// obterCotacao realiza a requisição HTTP para buscar a cotação do dólar
func obterCotacao(ctx context.Context) (*Cotacao, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlCotacao, nil)
	if err != nil {
		return nil, fmt.Errorf("Falha ao criar a requisição: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("Timeout na requisição")
		}
		return nil, fmt.Errorf("Erro ao fazer a requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Resposta com status inesperado: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Erro ao ler o corpo da resposta: %w", err)
	}

	var cotacao Cotacao
	if err := json.Unmarshal(body, &cotacao); err != nil {
		return nil, fmt.Errorf("Falha ao decodificar JSON: %w", err)
	}

	return &cotacao, nil
}

// salvarCotacao salva a cotação em um arquivo de texto
func salvarCotacao(cotacao *Cotacao) error {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		return fmt.Errorf("Erro ao criar o arquivo: %w", err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "Dólar: %.2f\n", cotacao.Valor)
	if err != nil {
		return fmt.Errorf("Erro ao escrever no arquivo: %w", err)
	}

	return nil
}
