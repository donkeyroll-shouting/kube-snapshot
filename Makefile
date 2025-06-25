# Define variables
GO_FILES := $(shell find . -type f -name "*.go")

# Recipe to compile the main package
build:
	go build -o _out/ksnap cmd/ksnap/main.go

# Recipe to run gofmt on all Go files
fmt:
	gofmt -w $(GO_FILES)