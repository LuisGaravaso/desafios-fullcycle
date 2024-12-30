# Etapa 1: Build
FROM golang:1.21 AS builder

# Define o diretório de trabalho no container
WORKDIR /app

# Copia o código fonte
COPY main.go .

# Compila o binário
RUN go build -o fullcycle main.go

# Etapa 2: Imagem mínima
FROM scratch

# Copia o binário da etapa de build
COPY --from=builder /app/fullcycle /fullcycle

# Define o comando padrão para executar o binário
ENTRYPOINT ["/fullcycle"]
