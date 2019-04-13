.PHONY: deps test

deps:
	@GO111MODULE=on go mod download

test: deps
	@go test ./...
