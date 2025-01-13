# Usa una imagen base de Go
FROM golang:1.21.3-alpine

# Instala las dependencias necesarias para CGO
RUN apk add --no-cache gcc musl-dev

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el archivo go.mod y go.sum y descarga las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el código fuente de la aplicación
COPY . .

# Habilita CGO y compila la aplicación
ENV CGO_ENABLED=1
RUN go build -o main ./cmd

# Expone el puerto en el que la aplicación se ejecutará
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
