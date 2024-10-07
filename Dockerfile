# Usar a imagem base do Go
FROM golang:latest

# Instalar as dependências para SSL e PKG Config no Debian/Ubuntu
RUN apt-get update && apt-get install -y \
    bash \
    gcc \
    musl-dev \
    libssl-dev \
    pkg-config \
    build-essential

# Definir o diretório de trabalho no container
WORKDIR /app

# Copiar o go.mod para o container (sem o go.sum, pois ele pode não existir)
COPY go.mod ./

# Baixar as dependências e gerar o go.sum
RUN go mod download
RUN go mod tidy  # Isso irá gerar o go.sum no container

# Copiar o restante do código-fonte para o container
COPY ./src .

# Instalar air para hot-reload
RUN go install github.com/air-verse/air@latest

# Corrigir o nome do pacote Tarantool
RUN go get github.com/tarantool/go-tarantool

# Copiar o arquivo .air.toml para o container
COPY ./.air.toml .

# Expor a porta da aplicação
EXPOSE 8080

# Comando para rodar a aplicação com hot-reload
CMD ["air", "-c", ".air.toml"]
