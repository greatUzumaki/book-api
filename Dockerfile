# syntax=docker/dockerfile:1

# Сборка
FROM golang:alpine AS builder

ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o server cmd/server/main.go

# Деплой
FROM scratch AS runner

WORKDIR /app

ENV DB_HOST=postgres_container
ENV GIN_MODE=release

COPY --from=builder /build/.env .
COPY --from=builder /build/server .

EXPOSE 8080

ENTRYPOINT ["./server"]