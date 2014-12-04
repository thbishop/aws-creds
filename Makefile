all: clean test build

binaries: clean fmt test
	@script/build_binaries.sh

build:
	@echo "==> Compiling source code."
	@godep go build -v -o ./bin/aws-creds ./aws-creds

clean:
	@echo "==> Cleaning up previous builds."
	@rm -rf bin/aws-creds

deps:
	@echo "==> Downloading dependencies."
	@godep save -r ./aws-creds/...

fmt:
	@echo "==> Formatting source code."
	@goimports -w ./aws-creds

test: fmt vet
	@echo "==> Running tests."
	@godep go test -cover ./aws-creds/...

vet:
	@godep go vet ./aws-creds/...

help:
	@echo "binaries\tcreate binaries"
	@echo "build\t\tbuild the code"
	@echo "clean\t\tremove previous builds"
	@echo "deps\t\tdownload dependencies"
	@echo "fmt\t\tformat the code"
	@echo "test\t\ttest the code"
	@echo "vet\t\tvet the code"
	@echo ""
	@echo "default will test, format, and build the code"

.PNONY: all clean deps fmt help test
