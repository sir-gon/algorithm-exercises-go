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

# DOCKER
DOCKER_COMPOSE=docker compose

ifeq ($(BRUTEFORCE),1)
  add_bruteforce=-tags bruteforce
endif
ifeq ($(BRUTEFORCE),TRUE)
  add_bruteforce=-tags bruteforce
endif

# Package Manager
GO=go
GOTEST=$(GO) test -count=1 -v $(add_bruteforce)
GOCOVER=$(GO) tool cover
PKG_LIST=$(go list ./... | grep -v /vendor/ | tr '\n' ' '| xargs echo -n)

# DOCKER
BUILDKIT_PROGRESS=plain
CGO_ENABLED=0

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
	$(GO) mod download -x
	$(GO) mod verify
	@echo "################################################################################"

lint/markdown:
	markdownlint '**/*.md' --ignore node_modules && echo '✔  Your code looks good.'

lint/yaml:
	yamllint --stric . && echo '✔  Your code looks good.'

lint: lint/markdown lint/yaml test/styling test/static

test/static: dependencies
	$(GO) vet -v ./...
	golangci-lint --verbose run ./...

test/styling: dependencies
	gofmt -l . && echo '✔  Your code looks good.'

coverage.out: env dependencies
	$(GOTEST) -v -covermode=atomic -coverprofile="coverage.out" ./exercises/...

test: env dependencies coverage.out
	$(GOCOVER) -func=coverage.out

coverage: test

coverage/html: coverage.out
	$(GOCOVER) -html=coverage.out -o ./coverage/coverage.html

outdated:
	$(GO) list -m -u all

update: dependencies outdated
	$(GO) get -u all
	$(GO) get -t -u ./...
	$(GO) mod tidy

upgrade: update

clean:
	$(GO) clean -testcache
	rm -vfr ./coverage
	mkdir -p ./coverage
	touch ./coverage/.gitkeep

build: env dependencies
	$(GO) build -v -o bin/ ./...

compose/build: env
	${DOCKER_COMPOSE} --profile lint build
	${DOCKER_COMPOSE} --profile testing build
	${DOCKER_COMPOSE} --profile production build

compose/rebuild: env
	${DOCKER_COMPOSE} --profile lint build --no-cache
	${DOCKER_COMPOSE} --profile testing build --no-cache
	${DOCKER_COMPOSE} --profile production build

compose/lint/markdown: compose/build
	${DOCKER_COMPOSE} --profile lint run --rm algorithm-exercises-go-lint make lint/markdown

compose/lint/yaml: compose/build
	${DOCKER_COMPOSE} --profile lint run --rm algorithm-exercises-go-lint make lint/yaml

compose/test/styling: compose/build
	${DOCKER_COMPOSE} --profile lint run --rm algorithm-exercises-go-lint make test/styling

compose/test/static: compose/build
	${DOCKER_COMPOSE} --profile lint run --rm algorithm-exercises-go-lint make test/static

compose/lint: compose/lint/markdown compose/lint/yaml compose/test/styling compose/test/static

compose/test: compose/build
	${DOCKER_COMPOSE} --profile testing run --rm algorithm-exercises-go-test make test

compose/run: compose/build
	${DOCKER_COMPOSE} --profile production run --rm algorithm-exercises-go

all: test coverage

run:
	ls -alh
