# syntax=docker/dockerfile:1

FROM golang:1.25 AS builder
WORKDIR /app

# чтобы make работал
RUN apt-get update && apt-get install -y --no-install-recommends make \
 && rm -rf /var/lib/apt/lists/*

# зависимости
COPY go.mod go.sum ./
RUN go mod download

# исходники
COPY . .

# СБОРКА через твой Makefile -> положит бинарь в build/app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build

# Рантайм-слой
FROM ubuntu:22.04
WORKDIR /app
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
 && rm -rf /var/lib/apt/lists/*

# копируем готовый бинарь
COPY --from=builder /app/build/app /app/app

EXPOSE 8080
ENTRYPOINT ["/app/app"]
