FROM golang:1.13

RUN mkdir -p /go/src/github.com/taniwhy/mochi-match-rest
WORKDIR /go/src/github.com/taniwhy/mochi-match-rest

ADD . /go/src/github.com/taniwhy/mochi-match-rest

EXPOSE 8000

ENV GO111MODULE=on