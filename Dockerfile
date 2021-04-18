ARG GO_VERSION=1.14

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR $GOPATH/src/github.com/crisaltmann/fundament-stock-api

COPY . .

#FROM alpine:latest
#RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

##RUN go get -d -v ./...
##RUN go install -v ./...

RUN go mod download
RUN go install -v ./...

EXPOSE 8080

ENTRYPOINT ["fundament-stock-api"]