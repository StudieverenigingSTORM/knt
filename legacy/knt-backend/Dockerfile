# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /app

COPY ./go.mod ./go.sum ./main.go ./
COPY ./internal ./internal

RUN go mod download

RUN GOOS=linux go build -o /knt .
