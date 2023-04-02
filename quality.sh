#!/bin/sh
echo "Running quality check"

echo "starting quality check."
echo "running go fmt"
go fmt ./...
echo "do a go mod tidy"
go mod tidy
echo "running go generate"
go generate ./...
echo "running golangci-lint"
golangci-lint run ./... -v --config .golangci.yml
echo "I don finish, all izz well?"
echo "quality check done :: please confirm !!! No Error?"
