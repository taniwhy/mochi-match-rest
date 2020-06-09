# Go パラメータ
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=main
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
run:
	cd db/
	rm dbconf.yml
	go get bitbucket.org/liamstask/goose/cmd/goose
	cp $GOPATH/src/bitbucket.org/liamstask/goose/db-sample/dbconf.yml ./
	goose -env environment_variable_config up
	$(GOBUILD) -o $(BINARY_NAME) -v ./main.go
	./$(BINARY_NAME)
hot:
	realize start
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
