## Run tests
test:
	@go test -race ./...

## Run tests with coverage
test-cover:
	@go test -coverprofile=cover.out -race ./...

## Build binary
build:
	@go build ./...

lint:
	@go fmt ./... && go vet ./...
