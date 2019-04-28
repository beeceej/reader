FROM golang:1.12-alpine

VOLUME /app

# Install the base we need for our Go tooling
RUN apk add --update bash ca-certificates curl gcc git musl-dev mysql-client xz

# By default, any command we run that's relative should be run from the
# builtin repository mount point.
RUN env
WORKDIR /app