GOPATH=$(shell go env GOPATH)

.PHONY: all
all: test lint

.PHONY: bin-deps
bin-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest


.PHONY: lint	
lint:
	$(GOPATH)/bin/golangci-lint run

.PHONY: test
test: 
	go test ./...