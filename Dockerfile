ARG GO_VERSION=1.15

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o .

EXPOSE 80
CMD ["./fundament-stock-api"]