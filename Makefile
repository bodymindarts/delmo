EXTERNAL_TOOLS=\
	       github.com/kardianos/govendor \
	       github.com/mitchellh/gox \

all: test dev

dev:
	@DELMO_DEV=1 sh -c "'$(PWD)/scripts/build.sh'"

test:
	scripts/test.sh

build:
	scripts/build.sh

bootstrap:
	@for tool in  $(EXTERNAL_TOOLS) ; do \
		echo "Installing $$tool" ; \
		go get $$tool; \
	done

PHONY: all test
