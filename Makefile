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
	$(GOBUILD) -o $(BINARY_NAME) -v ./main.go
	./$(BINARY_NAME)
test:
	$(GOTEST) -v ./...
deploy:
	$(GOBUILD) -o $(BINARY_NAME) -v ./main.go
	./$(BINARY_NAME)
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

server {
    listen       80 default_server;
    server_name drone.mochi-match.work;

    location / {
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
		proxy_set_header Connection "upgrade";
		roxy_set_header Host $host;

        proxy_pass http://drone_server;
        }
}


server {
    listen       80 default_server;
    server_name api.mochi-match.work;

    location / {
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
		proxy_set_header Connection "upgrade";
		roxy_set_header Host $host;

        proxy_pass http://app;
        }
}