FROM golang:latest

WORKDIR /go/src/github.com/taniwhy/mochi-match-rest

COPY . .

EXPOSE 8080

CMD [ "make", "run" ]