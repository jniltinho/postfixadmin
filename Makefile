GOCMD ?= go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOGET = $(GOCMD) get
BINARY_NAME = postfixadmin


default: get format generate build


get:
	@go mod tidy

generate:
	@templ generate


format:
	@templ fmt .

build:
	CGO_ENABLED=0 $(GOBUILD) -o $(BINARY_NAME) -v -ldflags="-s -w"


clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)


run:
	./$(BINARY_NAME)