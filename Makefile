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
	revive -config revive.toml -formatter friendly ./...
	go-critic check ./...

.PHONY: build test fmt lint