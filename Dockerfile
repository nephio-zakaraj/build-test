FROM golang:1.20.2-alpine3.17 AS builder

COPY . /build-test

RUN apk update && \
    apk add --no-cache make

WORKDIR /build-test

RUN go build -o bin/nmath

