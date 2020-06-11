FROM "golang"

WORKDIR /go/src/github.com/taniwhy/mochi-match-rest

COPY . .

EXPOSE 8000