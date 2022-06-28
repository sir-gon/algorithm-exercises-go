GO=go
GOTEST=$(GO) test
GOCOVER=$(GO) tool cover
PKG_LIST=$(go list ./... | grep -v /vendor/ | tr '\n' ' '| xargs echo -n)

.MAIN: test/coverage
.PHONY: all clean coverage dependencies help list test

help: list

list:
	@LC_ALL=C $(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

dependencies:
	$(GO) mod download

coverage/c.out: dependencies
	$(GOTEST) -v -coverprofile="coverage/c.out" ./...

test: coverage/c.out

coverage: coverage/c.out
	$(GOCOVER) -func=coverage/c.out
	$(GOCOVER) -html=coverage/c.out -o ./coverage/coverage.html

clean:
	$(GO) clean -testcache
	rm -vfr ./coverage
	mkdir -p ./coverage
	touch ./coverage/.gitkeep

all: test coverage
