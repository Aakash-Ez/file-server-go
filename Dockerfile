# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN mkdir -p temp

RUN go build -o docker-go-fileserver

CMD './docker-go-fileserver'