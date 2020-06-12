FROM golang:1.13-nanoserver-1809

ENV DOCKERIZE_VERSION v0.6.1
RUN apk add --no-cache openssl \
    && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

RUN mkdir -p /go/src/github.com/taniwhy/mochi-match-rest
WORKDIR /go/src/github.com/taniwhy/mochi-match-rest

ADD . /go/src/github.com/taniwhy/mochi-match-rest

EXPOSE 8000

ENV GO111MODULE=on