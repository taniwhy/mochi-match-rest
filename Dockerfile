FROM "golang"

WORKDIR /go/src/github.com/taniwhy/mochi-match-rest

COPY . .

EXPOSE 8080

CMD bash -c "go build && go run main.go"