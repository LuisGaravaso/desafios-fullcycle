
# Desafio Pós-Graduação: 📦 Consulta de CEP com Múltiplas APIs
Este projeto foi desenvolvido como parte dos requisitos de um desafio acadêmico da pós-graduação Go Expert da Full Cycle.

## Requisitos do Projeto
Implementar os ensinamentos das aulas de Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

https://brasilapi.com.br/api/cep/v1/ + cep

http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

Além dos desafios da Pós, eu me desafiei a colocar mais propriedades:

- Consultar multiplos CEPs de uma vez

- Permitem diferentes outputs de dados com as flags `-json`, `-text` e `raw`

## 📋 Funcionalidades

- **Consulta Simultânea**: Faz requisições em paralelo para **ViaCEP** e **BrasilAPI**.
- **Validação de CEP**: Verifica se o CEP fornecido está em um formato válido.
- **Formatos de Saída**: Suporta saída em:
  - JSON (`-json`)
  - Texto (`-text`)
  - RAW (resposta crua da API) (`-raw`)
- **Tratamento de Erros**: Informa quando o CEP é inválido ou quando ocorre um timeout.

## 🚀 Como Executar o Projeto

### Pré-requisitos

- **Go** instalado em sua máquina (versão 1.22.5 ou superior).

Para verificar a instalação do Go:

```bash
go version
```

### Clone o Repositório e Navegue para a pasta
```bash
git clone https://github.com/LuisGaravaso/desafios-fullcycle.git
cd desafios-fullcycle/Goexpert\ -\ Multithreading/
```

### Compilar e Executar

Execute o programa passando os CEPs como argumentos:

```bash
go run main.go <CEP> [<CEP> ...] [opções de saída]
```

#### Exemplos de Execução

1. **Consulta com saída em JSON (padrão):**

   ```bash
   go run main.go 01001-000 22041-001
   ```

2. **Consulta com saída em texto:**

   ```bash
   go run main.go -text 01001-000
   ```

3. **Consulta com saída crua da API:**

   ```bash
   go run main.go 01001-000 -raw
   ```

#### Exemplo de Saída em Texto

```
Praça da Sé, Bairro Sé 
São Paulo-SP 
01001-000 
Dados Obtidos de viacep 
------------
```

#### Exemplo de Saída em JSON

```json
{
  "cep": "01001-000",
  "logradouro": "Praça da Sé",
  "bairro": "Sé",
  "cidade": "São Paulo",
  "estado": "SP",
  "service": "viacep"
}
```

## 📂 Estrutura do Projeto

```
meu-projeto/
│-- main.go
│-- go.mod
│
├── models/
│   └── cep.go          # Structs e interfaces para mensagens de CEP
│
├── services/
│   └── request.go      # Função para realizar requisições HTTP
│
├── utils/
│   └── argsparsers.go      # Funções auxiliaraes para validação dos Argumentos passados no Terminal
│   └── cepparser.go      # Funções auxiliaraes para validação dos CEPs
│
└── handlers/
    └── process.go      # Função para processar CEPs e formatar a saída
```

## ⚙️ Tecnologias Utilizadas

- **Go**: Linguagem de programação principal.
- **APIs Utilizadas**:
  - [ViaCEP](https://viacep.com.br/)
  - [BrasilAPI](https://brasilapi.com.br/)

## 🛠️ Principais Pacotes

- **`net/http`**: Para fazer requisições HTTP.
- **`encoding/json`**: Para decodificar respostas em JSON.
- **`sync`**: Para gerenciar concorrência com `WaitGroup`.
- **`context`**: Para gerenciar timeout nas requisições.

## ❗ Tratamento de Erros

- **CEP Inválido**: O programa identifica e exibe os CEPs inválidos.
- **Timeout**: Se nenhuma API responder em 1 segundo, é exibida uma mensagem de timeout.

## 📝 Licença

Este projeto é licenciado sob a **MIT License**. Sinta-se à vontade para usar, modificar e compartilhar!
