GOPATH=$(shell go env GOPATH)

.PHONY: all
all: test lint

.PHONY: bin-deps
bin-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
	go get golang.org/x/tools/cmd/cover@latest


.PHONY: lint	
lint:
	$(GOPATH)/bin/golangci-lint run

.PHONY: test
test: 
	go clean -testcache
	go test ./...

.PHONY: coverage
coverage:
	go test -v -coverprofile=cover.out ./...
	go tool cover -html=cover.out