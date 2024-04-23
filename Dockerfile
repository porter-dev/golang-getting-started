# syntax=docker/dockerfile:1.1.7-experimental

# Environment to build manager binary
FROM golang:1.15.6-alpine3.12 as build
WORKDIR /gin

RUN apk update && apk add gcc musl-dev

COPY . .

RUN go mod download

RUN go build -ldflags '-w -s' -a -o ./bin/porter-gin-example .

# Deployment environment
# ----------------------
FROM alpine:3.12
WORKDIR /gin

RUN apk update && apk add git curl

COPY --from=build /gin/bin/porter-gin-example /usr/bin/

ENTRYPOINT [ "porter-gin-example" ]

