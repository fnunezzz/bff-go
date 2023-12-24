MAIN_PACKAGE_PATH := ./cmd/bff
BINARY_NAME=app

test:
	@go test -v ./...

dev:
	go run cmd/bff/main.go

build:
	@go build -v -o=bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}
	
build-docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o=bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}