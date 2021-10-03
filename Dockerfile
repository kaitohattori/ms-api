FROM golang:1.14-alpine3.11 AS build

WORKDIR /go/src/ms-api

COPY /app ./app
COPY /config ./config
COPY /docs ./docs
COPY /scripts ./scripts
COPY go.mod ./
COPY go.sum ./
COPY main.go ./

ENV GO111MODULE=on

RUN go mod download

RUN go build main.go

ENTRYPOINT go run main.go
