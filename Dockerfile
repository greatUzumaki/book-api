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

WORKDIR /


ARG PORT=https://cdn.communism.ru
ENV PORT=$PORT

ENV DB_HOST=postgres_container
ENV PORT=8080
ENV DB_LOGIN=postgres
ENV DB_PASSWORD=pswd1234
ENV DB_NAME=postgres
ENV DB_PORT=5432

ENV GIN_MODE=release

COPY --from=builder /build/.env /
COPY --from=builder /build/server /server

EXPOSE $PORT

ENTRYPOINT ["/server"]