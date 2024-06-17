# Etapa de build
FROM golang:1.22.4 AS builder

WORKDIR /app

# Copie os arquivos go.mod e go.sum e baixe as dependências
COPY . ./
RUN go mod tidy
RUN go get github.com/go-playground/validator/v10@v10.20.0

# Copie o código fonte
COPY . .

# Compile a aplicação
RUN CGO_ENABLED=1 GOOS=linux go build -o backend-golang

# Etapa final
FROM debian:latest

WORKDIR /app

# Copie o binário compilado da etapa de build
COPY --from=builder /app/backend-golang .

# Exponha a porta que a aplicação vai rodar
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./backend-golang"]
