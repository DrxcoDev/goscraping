# Usar una imagen de Go para compilar la aplicación
FROM golang:1.22 AS builder

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar go.mod y go.sum y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto de la aplicación
COPY . .

# Compilar la aplicación
RUN go build -o /build/scraper

# Usar una imagen de Alpine para ejecutar la aplicación
FROM alpine:latest

# Copiar el binario compilado desde la imagen de builder
COPY --from=builder /build/scraper /usr/local/bin/scraper

# Establecer el comando por defecto
CMD ["scraper"]
