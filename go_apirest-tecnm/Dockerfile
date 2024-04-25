# Utiliza una imagen base de Go
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el código fuente de tu API al directorio de trabajo del contenedor
COPY . .

# Compila el código de tu API
RUN go build -o main .

# Expone el puerto en el que tu API escucha
EXPOSE 8080

# Comando por defecto para ejecutar tu API
CMD ["./main"]
