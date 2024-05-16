#build stage
FROM golang:1.22.2-alpine3.19 AS builder
RUN apk add --no-cache git

WORKDIR /app

COPY ["go.mod", "go.sum", "./"]
RUN go mod download -x

COPY . .

RUN go build \ 
    -ldflags="-s -w" \
    -o app -v .

#final stage
FROM alpine:3.19
LABEL Name=dockerization

RUN apk update
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app .

ENTRYPOINT ["./app"]

EXPOSE 3000

#Crear contenedor de la api para que se pueda conectar a otro contenedor. 
#docker run -d --name restapi-1 -p 3000:3000 --env POSTGRES_HOST=some-postgres --env POSTGRES_PORT=5432 
#--env POSTGRES_USER=david --env POSTGRES_PASSWORD=password --env POSTGRES_DB=db_apitecnm apitecnm:v1