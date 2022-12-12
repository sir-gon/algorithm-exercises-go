FROM golang:1.19.3-alpine3.16 as base

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

FROM builder as testing

ENV LOG_LEVEL=INFO
ENV BRUTEFORCE=FALSE

RUN adduser -D worker
RUN mkdir -p /app
RUN chown worker:worker /app

WORKDIR /app

RUN ls -alh

USER worker
CMD ["make", "test"]
