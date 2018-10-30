version=0.1.0

.PHONY: all

all:
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@echo "  build         - build the source code"
	@echo "  lint          - lint the source code"
	@echo "  test          - test the source code"
	@echo "  fmt           - format the code with gofmt"
	@echo "  install       - install dependencies"

lint:
	@go vet $(go list ./...)
	@go list ./... | grep -v /vendor/ | xargs -L1 golint

test:
	@ginkgo -r

fmt:
	@go fmt $(go list ./...)

build: lint
	@go build $(go list ./...)

install:
	@go get -u github.com/golang/lint/golint
	@go get -u github.com/onsi/ginkgo/ginkgo
	@go mod vendor
