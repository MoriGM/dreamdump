PKGNAME=dreamdump
GO_CRITIC_FLAGS=-@ifElseChain.minThreshold=4 -enable='octalLiteral,yodaStyleExpr,zeroByteRepeat,badSorting,builtinShadow,commentFormatting'

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
	go-critic check ${GO_CRITIC_FLAGS} ./...

install:
	go install .

.PHONY: build test fmt lint install