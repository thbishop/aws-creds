base_dir = $(shell pwd)
gopath = "$(base_dir)/vendor:$(GOPATH)"

all: check-gopath clean test fmt build

build:
	@echo "==> Compiling source code."
	@env GOPATH=$(gopath) go build -v -o ./bin/aws-creds ./aws-creds

test: check-gopath
	@echo "==> Running tests."
	@env GOPATH=$(gopath) go test -cover ./aws-creds/...

deps: check-gopath
	@echo "==> Downloading dependencies."
	@env GOPATH=$(gopath) go get -d -v ./aws-creds/...

	@echo "==> Removing .git and .bzr from vendor."
	@find ./vendor -type d -name .git | xargs rm -rf
	@find ./vendor -type d -name .bzr | xargs rm -rf
	@find ./vendor -type d -name .hg | xargs rm -rf

fmt:
	@echo "==> Formatting source code."
	@gofmt -w ./aws-creds

clean:
	@echo "==> Cleaning up previous builds."
	@rm -rf bin/aws-creds

help:
	@echo "build\t\tbuild the code"
	@echo "clean\t\tremove previous builds"
	@echo "deps\t\tdownload dependencies"
	@echo "fmt\t\tformat the code"
	@echo "test\t\ttest the code"
	@echo ""
	@echo "default will test, format, and build the code"

check-gopath:
ifndef GOPATH
  $(error GOPATH is undefined)
endif

.PNONY: all clean deps fmt help test
