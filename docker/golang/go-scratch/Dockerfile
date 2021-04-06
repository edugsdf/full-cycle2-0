
#----------------------------------------------------------------
# 01) Criar imagem (gorda) com todas as dependencias
FROM golang:1.16.2-alpine3.13 as builder

# Diretorio de trabalho
RUN mkdir -p /app
WORKDIR /app

### arquivos de dependencias
COPY go.mod .

### Configurando proxy para downloads
ENV GOPROXY https://proxy.golang.org,direct

### baixando dependencias modulos
RUN go mod download

### copiando local para container
COPY . .

### Desativando o CGO
### caso contrario o aplicativo pode não conseguir iniciar
ENV CGO_ENABLED=0

### Construindo a aplicação GO para Linux
RUN GOOS=linux go build ./app.go

#----------------------------------------------------------------
# 02) Definir qual imagem executar
FROM scratch

### Diretório de trabalho
WORKDIR /app

### Copiar o binário da imagem da etapa 01
COPY --from=builder /app/app .

### Executa binário da aplicação
CMD ["/app/app"]
