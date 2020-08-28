FROM golang:1.15.0-buster

ENV APP_ROOT=/go/src/github.com/tkuchiki/go-workers-test
ENV GO111MODULE=on

RUN mkdir -p $APP_ROOT

WORKDIR $APP_ROOT
ADD . $APP_ROOT/

RUN go mod download
