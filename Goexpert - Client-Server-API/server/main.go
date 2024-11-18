package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "modernc.org/sqlite"
)

type CotacaoIn struct {
	USDBRL struct {
		Code   string  `json:"code"`
		Codein string  `json:"codein"`
		Name   string  `json:"name"`
		Bid    float64 `json:"bid,string"`
	} `json:"USDBRL"`
}

type CotacaoOut struct {
	Valor float64 `json:"valor"`
}

func main() {
	db, err := connectToDatabase()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		getDolarValueHandler(w, r, db)
	})

	log.Println("Servidor iniciado na porta 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

func getDolarValueHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	cotacao, err := requestDolarValue()
	if err != nil {
		http.Error(w, "Erro ao obter cotação", http.StatusInternalServerError)
		log.Printf("Erro ao obter cotação: %v", err)
		return
	}

	cotacaoOut := CotacaoOut{Valor: cotacao.USDBRL.Bid}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cotacaoOut); err != nil {
		http.Error(w, "Erro ao codificar resposta JSON", http.StatusInternalServerError)
		log.Printf("Erro ao codificar JSON: %v", err)
		return
	}

	if err := insertToDatabase(db, cotacao); err != nil {
		http.Error(w, "Erro ao salvar cotação", http.StatusInternalServerError)
		log.Printf("Erro ao salvar cotação no banco de dados: %v", err)
	}
}

// connectToDatabase cria a Conexão com o Banco SQLite
func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./cotacoes.db")
	if err != nil {
		return nil, err
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS cotacoes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT NOT NULL,
		code_in TEXT NOT NULL,
		name TEXT NOT NULL,
		value REAL NOT NULL,
		timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
	);`
	if _, err := db.Exec(createTable); err != nil {
		return nil, err
	}

	return db, nil
}

// insertToDatabase insere a cotação do dólar no Banco de Dados
func insertToDatabase(db *sql.DB, cotacao CotacaoIn) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	insertSTMT := "INSERT INTO cotacoes (code, code_in, name, value) VALUES(?,?,?,?)"
	_, err := db.ExecContext(ctx, insertSTMT, cotacao.USDBRL.Code, cotacao.USDBRL.Codein, cotacao.USDBRL.Name, cotacao.USDBRL.Bid)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Fatalf("Timeout ao inserir no Banco de Dados: %w", err)
		}
	}
	return err
}

// requestDolarValue faz a requisição na API para pegar a cotação do Dólar
func requestDolarValue() (CotacaoIn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return CotacaoIn{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("Tempo de Request Excedido")
		}
		return CotacaoIn{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return CotacaoIn{}, fmt.Errorf("API retornou status: %d", resp.StatusCode)
	}

	var cotacaoIn CotacaoIn
	if err := json.NewDecoder(resp.Body).Decode(&cotacaoIn); err != nil {
		return CotacaoIn{}, err
	}

	log.Printf("Cotação do Dólar Enviada com Sucesso")
	return cotacaoIn, nil
}
