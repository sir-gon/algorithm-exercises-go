###############################################################################
FROM golang:1.22.3-alpine3.19 AS base

ENV CGO_ENABLED 0
RUN apk add --update --no-cache make

ENV WORKDIR=/app
WORKDIR ${WORKDIR}

###############################################################################
FROM node:22.1.0-alpine3.19 AS lint

ENV WORKDIR=/app
WORKDIR ${WORKDIR}

COPY ./docs ${WORKDIR}/docs
RUN apk add --update --no-cache make
RUN npm install -g markdownlint-cli

###############################################################################
FROM base AS development

ENV BINDIR /usr/local/bin

# In alpine linux (as it does not come with curl by default)
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.53.3

###############################################################################
FROM development AS builder

COPY ./src ${WORKDIR}/src
COPY ./go.mod ${WORKDIR}/
COPY ./go.sum ${WORKDIR}/
COPY ./Makefile ${WORKDIR}/

RUN make dependencies

###############################################################################
### In testing stage, can't use USER, due permissions issue
## in github actions environment:
##
## https://docs.github.com/en/actions/creating-actions/dockerfile-support-for-github-actions
##
FROM builder AS testing

ENV LOG_LEVEL=INFO
ENV BRUTEFORCE=FALSE

WORKDIR /app

RUN ls -alh

CMD ["make", "test"]

###############################################################################
### In production stage
## in the production phase, "good practices" such as
## WORKSPACE and USER are maintained
##
FROM builder AS production

ENV LOG_LEVEL=INFO
ENV BRUTEFORCE=FALSE

RUN adduser -D worker
RUN mkdir -p /app
RUN chown worker:worker /app

WORKDIR /app

RUN ls -alh

USER worker
CMD ["make", "test"]
