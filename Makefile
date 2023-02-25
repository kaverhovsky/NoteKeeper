# go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMODDOWNLOAD=$(GOCMD) mod tidy

APP_ENTRYPOINT=cmd/note_service.go
BINARY_NAME=note_service
BUILD_FOLDER=build

.PHONY: all build clean deps

all: clean build
build:
	mkdir -p $(BUILD_FOLDER)
	cp -n .env.example $(BUILD_FOLDER)/.env
	$(GOBUILD) -o $(BUILD_FOLDER)/$(BINARY_NAME) -v $(APP_ENTRYPOINT)
deps:
	$(GOMODOWNLOAD)
clean:
	$(GOCLEAN)
	rm -rf $(BUILD_FOLDER)/*
test:
	$(GOTEST) -v ./...
start:
	cp -n .env.example $(BUILD_FOLDER)/.env
	./$(BUILD_FOLDER)/$(BINARY_NAME) -c $(BUILD_FOLDER)/.env
run: clean build start
