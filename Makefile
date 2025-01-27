APP_NAME=go-clip

build:
	go build -o build/$(APP_NAME) bin/main.go

.PHONY: build

publish:
	goreleaser release --snapshot --clean

.PHONY: publish

clean:
	@rm -rf dist
	@rm -rf build

.PHONY: clean

credits:
	@gocredits > CREDITS