## REFERENCES:
## (1) Passing environment variable with fallback value to Makefile:
##    https://stackoverflow.com/a/70772707/6366150
## (2) Export environment variables inside "make environment"
##		https://stackoverflow.com/a/49524393/6366150
## (3) Uppercase to lowercase and vice versa
##		https://community.unix.com/t/uppercase-to-lowercase-and-vice-versa/285278/6
## (4) How do I trim leading and trailing whitespace from each line of some output?
## 		https://unix.stackexchange.com/a/279222/233927
############################################################################

## (1) ## Allowed values: info | warn | error | debug
LOG_LEVEL ?= info
## (3) (4)
LOG_LEVEL :=$(shell echo '${LOG_LEVEL}'| tr '[:lower:]' '[:upper:]'| tr -d '[:blank:]')

## (1) ## Allowed values: true | false
BRUTEFORCE ?= false
## (3) (4)
BRUTEFORCE :=$(shell echo '${BRUTEFORCE}'| tr '[:lower:]' '[:upper:]'| tr -d '[:blank:]')

ifeq ($(BRUTEFORCE),1)
  add_bruteforce=-tags bruteforce
endif
ifeq ($(BRUTEFORCE),TRUE)
  add_bruteforce=-tags bruteforce
endif

# Package Manager
GO=go
GOTEST=$(GO) test $(add_bruteforce)
GOCOVER=$(GO) tool cover
PKG_LIST=$(go list ./... | grep -v /vendor/ | tr '\n' ' '| xargs echo -n)

# DOCKER
BUILDKIT_PROGRESS=plain

.MAIN: test/coverage
.PHONY: all clean coverage dependencies help list test
.EXPORT_ALL_VARIABLES: # (2)

help: list

list:
	@echo "Environment variables:"
	@echo "LOG_LEVEL = info | warn | error | debug"
	@echo "BRUTEFORCE = true | false"
	@echo ""
	@echo "Posible make commands:"
	@LC_ALL=C $(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

env:
	@echo "################################################################################"
	@echo "## Environment: ################################################################"
	@echo "################################################################################"
	@printenv | grep -E "LOG_LEVEL|BRUTEFORCE"
	@echo "################################################################################"

dependencies:
	@echo "################################################################################"
	@echo "## Dependencies: ###############################################################"
	@echo "################################################################################"
	$(GO) mod download
	@echo "################################################################################"

coverage/c.out: env dependencies
	$(GOTEST) -v -covermode=atomic -coverprofile="coverage/c.out" ./...

test: coverage/c.out

coverage: coverage/c.out
	$(GOCOVER) -func=coverage/c.out
	$(GOCOVER) -html=coverage/c.out -o ./coverage/coverage.html

clean:
	$(GO) clean -testcache
	rm -vfr ./coverage
	mkdir -p ./coverage
	touch ./coverage/.gitkeep

compose/build: env
	docker-compose --profile testing build

compose/rebuild: env
	docker-compose --profile testing build --no-cache

compose/run: compose/build
	docker-compose --profile testing run --rm algorithm-exercises-go make test

all: test coverage
