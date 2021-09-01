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

install-tools: fetch
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %
