PKGNAME=dreamdump

build: fmt test
	go build -v ./...
	go build -v .

test: fmt
	go test -v ./...

fmt:
	gofumpt -w .

lint: fmt
	golangci-lint run ./...
	staticcheck

.PHONY: build test fmt lint