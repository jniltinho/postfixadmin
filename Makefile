GOCMD ?= go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean -cache
GOGET = $(GOCMD) get
BINARY_NAME = postfixadmin


default: get tailwind templ build


get:
	@go mod tidy

tailwind:
	npx tailwindcss -i view/css/app.css -o public/static/styles.css

templ:
	@templ fmt view
	@templ generate view

build:
	CGO_ENABLED=0 $(GOBUILD) -o $(BINARY_NAME) -v -ldflags="-s -w"
	upx $(BINARY_NAME)


clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)


run:
	./$(BINARY_NAME) serve --config=local.toml



build-dev:
	CGO_ENABLED=0 $(GOBUILD) -tags=dev -o ./tmp/main .


dev: get tailwind templ build-dev
	./tmp/main serve --config=local.toml


install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss
	@npm install -D daisyui@latest
