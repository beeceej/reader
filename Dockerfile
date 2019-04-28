FROM golang:1.12-alpine

VOLUME /app

RUN apk add --update bash ca-certificates curl gcc git musl-dev mysql-client xz

RUN env
WORKDIR /app