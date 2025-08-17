PKGNAME=dreamdump

build:
	go build .

test:
	go test ./...

fmt:
	gofumpt -w .

lint:
	golangci-lint run ./...

.PHONY: build test fmt lint