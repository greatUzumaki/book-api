# syntax=docker/dockerfile:1

# Сборка
FROM golang:alpine AS builder

ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /app/hello . /hello.go

# Деплой
FROM scratch AS deploy

WORKDIR /

COPY --from=builder /godocker /godocker

EXPOSE 8080

ENTRYPOINT ["/godocker"]