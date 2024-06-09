APP_BINARY=kata-machina

.DEFAULT_GOAL := run

build: clean
	@go build -o bin/${APP_BINARY} cmd/main.go

run: build
	@./bin/${APP_BINARY} $(args)

.PHONY: clean
clean:
	@go clean
	@rm -f bin/${APP_BINARY}

format:
	@go fmt ./...

.PHONY: test
test:
	@go test ./dsa_test