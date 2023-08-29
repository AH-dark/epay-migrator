build:
	@go build -o bin/$(BINARY_NAME) -v

test:
	@go test -v ./...
