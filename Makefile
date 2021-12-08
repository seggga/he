LOCAL_BIN=$(CURDIR)/bin
GIT_COMMIT=$(shell git rev-list -1 HEAD)

# build the binary
.PHONY: build
build:
	mkdir -p $(LOCAL_BIN)
	go build -ldflags "-X main.CommitVer=$(GIT_COMMIT)" -o $(LOCAL_BIN)/sqlparser $(CURDIR)/cmd/sqlparser

# simple build and run the application
.PHONY: run
run:
	go run ./cmd/sqlparser/main.go

# run tests
.PHONY: test
test:
	go test -v ./...

# run linters 
.PHONY: lint
lint:
	golangci-lint run ./...

# generate pre-commit hooks accouding to .pre-commit-config.yaml
# .PHONY: pre-commit
# pre-commit:
# 	pre-commit install

.DEFAULT_GOAL := run