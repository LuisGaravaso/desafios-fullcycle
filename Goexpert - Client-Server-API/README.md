# Desafio Pós-Graduação: Client-Server API

Este projeto foi desenvolvido como parte dos requisitos de um desafio acadêmico da pós-graduação Go Expert da Full Cycle. 

O objetivo é implementar dois sistemas em Go (`client.go` e `server.go`) que utilizem contextos, webserver HTTP, banco de dados SQLite e manipulação de arquivos para processar e armazenar a cotação do dólar.  

## Requisitos do Projeto  

1. **Arquitetura do Sistema**  
   - **`client.go`**  
     - Envia uma requisição HTTP ao `server.go` para obter a cotação do dólar.  
     - Recebe apenas o valor atual da cotação (campo `bid` do JSON).  
     - Salva a cotação em um arquivo `cotacao.txt` no formato:  
       ```plaintext  
       Dólar: {valor}  
       ```  
     - Utiliza o package `context` com timeout máximo de **300ms** para processar a resposta.  
     - Loga erros caso o tempo de execução exceda o limite.  

   - **`server.go`**  
     - Consome a API externa: [AwesomeAPI](https://economia.awesomeapi.com.br/json/last/USD-BRL).  
     - Retorna a cotação para o cliente no endpoint `/cotacao` via JSON.  
     - Registra cada cotação no banco de dados SQLite.  
     - Utiliza os seguintes timeouts:  
       - Timeout máximo de **200ms** para chamar a API de cotação do dólar.  
       - Timeout máximo de **10ms** para persistir dados no banco.  
     - Loga erros em caso de timeouts ou falhas.  

2. **Especificações Técnicas**  
   - O servidor HTTP deve rodar na porta **8080**.  
   - A base de dados utilizada é SQLite para persistência das cotações.  
   - A API consumida para obter a cotação do dólar é:  
     ```
     https://economia.awesomeapi.com.br/json/last/USD-BRL  
     ```  

3. **Logs e Erros**  
   - Todos os contextos (`client`, API, e persistência no banco de dados) devem logar erros caso os tempos de execução sejam insuficientes.  

## Como Executar  

1. Clone este repositório:  
   ```bash  
   git clone https://github.com/LuisGaravaso/desafios-fullcycle.git
   cd Goexpert - Client-Server-API
   ```
   
2. Navegue para as pastas individuais e execute os programas.  

### Iniciar o Servidor  
No diretório `server`:  
```bash  
go run main.go  
```  

### Executar o Cliente  
No diretório `client`:  
```bash  
go run main.go  
```  

## Estrutura do Projeto  

```plaintext
📁 projeto-cotacao  
├── 📁 client  
│   └── main.go  
├── 📁 server
│   └── go.mod
│   └── go.sum  
│   └── main.go  
└── README.md  
```  

## Exemplo de Saída  

- **Arquivo `cotacao.txt`**:  
  ```plaintext  
  Dólar: 5.74  
  ```  

- **Resposta do servidor no endpoint `/cotacao`**:  
  ```json  
  {  
    "valor": "5.7473"  
  }  
  ```  

## Tecnologias Utilizadas  

- Linguagem: Go  
- Banco de Dados: SQLite  
- Ferramentas: `context`, `http`, `log`, `os`, `database/sql`  
```  


  
