PKGNAME=dreamdump

build: fmt
	go build -v ./...
	go build -v .

test: fmt
	go test -v ./...

fmt: lint
	gofumpt -w .

lint: fmt
	golangci-lint run ./...
	staticcheck
	errcheck

.PHONY: build test fmt lint