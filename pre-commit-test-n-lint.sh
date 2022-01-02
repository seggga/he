#!/bin/sh

# Turn on script stopping after the first error occured.
set -e

# Start testing
echo "start testing..."
go test -cover ./...

# Start linters.
echo "start linter..."
${HOME}/go/bin/golangci-lint run

echo "All checks successfully passed."