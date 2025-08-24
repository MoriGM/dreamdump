PKGNAME=dreamdump

build: test
	go build -v ./...
	go build -v .

test: fmt
	go test -v ./...

fmt: lint
	gofumpt -w .

lint:
	golangci-lint run ./...
	staticcheck
	errcheck

.PHONY: build test fmt lint