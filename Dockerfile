FROM golang:1.13

RUN mkdir -p /go/src/github.com/taniwhy/mochi-match-rest
WORKDIR /go/src/github.com/taniwhy/mochi-match-rest

ADD . /go/src/github.com/taniwhy/mochi-match-rest

ENV GO111MODULE=on

EXPOSE 8000

RUN go get github.com/pilu/fresh

CMD ["fresh"]