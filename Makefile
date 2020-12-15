# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=gim
BINARY_UNIX=$(BINARY_NAME)_unix
MAINPATH=main.go

all:build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAINPATH)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAINPATH)
	./$(BINARY_NAME)