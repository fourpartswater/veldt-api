version=0.1.0

.PHONY: all

NOVENDOR := $(shell glide novendor)

all:
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@echo "  build         - build the source code"
	@echo "  lint          - lint the source code"
	@echo "  test          - test the source code"
	@echo "  fmt           - format the code with gofmt"
	@echo ""
	@echo "  deps          - install tool dependencies"

lint:
	@go vet $(NOVENDOR)
	@go list ./... | grep -v /vendor/ | xargs -L1 golint

test:
	@go test $(NOVENDOR)

fmt:
	@go fmt $(NOVENDOR)

build: lint
	@go build $(NOVENDOR)

deps:
	@go get github.com/golang/lint/golint
	@go get github.com/Masterminds/glide
	@glide install
