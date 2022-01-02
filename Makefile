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
	go test -count=1 -cover ./...

# run linters 
.PHONY: lint
lint:
	golangci-lint run ./...

.DEFAULT_GOAL := run