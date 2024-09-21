FROM golang:1.17.8-alpine

RUN apk add --no-cache make git
RUN mkdir -p /go/src/github.com/abtreece/confd && \
  ln -s /go/src/github.com/abtreece/confd /app

WORKDIR /app
