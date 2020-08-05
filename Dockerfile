FROM golang:1.13

ENV DOCKERIZE_VERSION v0.6.1
ENV GO111MODULE=on

RUN go get github.com/rubenv/sql-migrate/...

RUN apt-get update && apt-get install -y wget \
    && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

RUN mkdir -p /go/src/github.com/taniwhy/mochi-match-rest
WORKDIR /go/src/github.com/taniwhy/mochi-match-rest
ADD . /go/src/github.com/taniwhy/mochi-match-rest

EXPOSE 8000
