FROM golang:1.18.3-alpine3.16 as base
ENV WORKDIR=/app
WORKDIR ${WORKDIR}

FROM base as development

FROM development as builder
RUN apk add --update --no-cache make

COPY ./src ${WORKDIR}/src
COPY ./go.mod ${WORKDIR}/
COPY ./Makefile ${WORKDIR}/

RUN make dependencies

FROM builder as testing

WORKDIR /app

RUN ls -alh

CMD ["make", "all"]
