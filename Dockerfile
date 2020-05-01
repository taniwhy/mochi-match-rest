FROM golang:latest

WORKDIR /go/src/github.com/taniwhy/mochi-match-rest

COPY . .

ENV GO111MODULE=on

EXPOSE 8080

CMD [ "make", "run" ]