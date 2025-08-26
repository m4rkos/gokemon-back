# Etapa 1: Build da aplicação
FROM golang:1.23-alpine AS builder

# Instala dependências essenciais
RUN apk add --no-cache git

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos go.mod e go.sum e baixa as dependências
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copia todo o restante da aplicação
COPY . .

# Compila a aplicação (output: main)
RUN go build -o main .

# Etapa 2: Imagem final (bem mais leve)
FROM alpine:latest

# Diretório de trabalho no container
WORKDIR /app

# Copia o binário da etapa anterior
COPY --from=builder /app/main .

# Expõe a porta da API (ajusta se tua app usa outra porta)
EXPOSE 8080

# Comando de entrada
CMD ["./main"]
