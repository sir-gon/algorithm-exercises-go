FROM golang:1.20.1-alpine3.16 as base

ENV CGO_ENABLED 0
RUN apk add --update --no-cache make

ENV WORKDIR=/app
WORKDIR ${WORKDIR}

FROM base as development

FROM development as builder

COPY ./src ${WORKDIR}/src
COPY ./go.mod ${WORKDIR}/
COPY ./go.sum ${WORKDIR}/
COPY ./Makefile ${WORKDIR}/

RUN make dependencies

### In testing stage, can't use USER, due permissions issue
## in github actions environment:
##
## https://docs.github.com/en/actions/creating-actions/dockerfile-support-for-github-actions
##
FROM builder as testing

ENV LOG_LEVEL=INFO
ENV BRUTEFORCE=FALSE

WORKDIR /app

RUN ls -alh

CMD ["make", "test"]

### In production stage
## in the production phase, "good practices" such as
## WORKSPACE and USER are maintained
##
FROM builder as production

ENV LOG_LEVEL=INFO
ENV BRUTEFORCE=FALSE

RUN adduser -D worker
RUN mkdir -p /app
RUN chown worker:worker /app

WORKDIR /app

RUN ls -alh

USER worker
CMD ["make", "test"]
