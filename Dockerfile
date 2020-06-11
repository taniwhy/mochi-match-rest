FROM "golang"

WORKDIR /go/src/github.com/taniwhy/mochi-match-rest

COPY . .

EXPOSE 8000

CMD go run main.go