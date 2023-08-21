run-dev:
	@go build .
	./sql-to-ts

build:
	@go build .

install-deps:
	@go mod tidy
