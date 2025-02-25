# Etapa de build
FROM golang:1.23.5

WORKDIR /app

# Copia os arquivos go.mod e go.sum e faz o download das dependências
COPY go.mod go.sum ./
RUN go mod tidy

# Copia o código-fonte da aplicação
COPY . .

# Compila a aplicação Go
RUN go build -o app ./cmd

# Expõe a porta 22 (SSH)
EXPOSE 22

# Comando para rodar o container
CMD ["/app/app"]
