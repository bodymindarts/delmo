PACKAGES = $(shell go list ./... | grep -v '/vendor/')
EXTERNAL_TOOLS=\
	       github.com/kardianos/govendor \
	       github.com/mitchellh/gox \
	       github.com/maxbrunsfeld/counterfeiter

all: test dev

dev: format
	@DELMO_DEV=1 sh -c "'$(PWD)/scripts/build.sh'"

test: generate
	scripts/test.sh

build:
	scripts/build.sh

format:
	@echo "--> Running go fmt"
	@go fmt $(PACKAGES)

generate:
	@echo "--> Running go generate"
	@go generate $(PACKAGES)

bootstrap:
	@for tool in  $(EXTERNAL_TOOLS) ; do \
		echo "Installing $$tool" ; \
		go get $$tool; \
	done

PHONY: all test build
