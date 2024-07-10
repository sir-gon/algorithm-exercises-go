###############################################################################
FROM golang:1.22.5-alpine3.20 AS init

ENV CGO_ENABLED 0
RUN apk add --update --no-cache make

ENV WORKDIR=/app
WORKDIR ${WORKDIR}

###############################################################################
FROM init AS base

ENV WORKDIR=/app
WORKDIR ${WORKDIR}

COPY ./go.mod ${WORKDIR}/
COPY ./go.sum ${WORKDIR}/
COPY ./Makefile ${WORKDIR}/
RUN make dependencies

###############################################################################
FROM base AS lint

ENV WORKDIR=/app
WORKDIR ${WORKDIR}

RUN apk add --update --no-cache make nodejs npm wget
RUN apk add --update --no-cache yamllint

RUN npm install -g --ignore-scripts markdownlint-cli

# golangci-lint
RUN wget --secure-protocol=TLSv1_2 --max-redirect=0 -O- -nv \
  https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
  sh -s -- -b $(go env GOPATH)/bin v1.59.1

# [!TIP] Use a bind-mount to "/app" to override following "copys"
# for lint and test against "current" sources in this stage

# YAML sources
COPY ./.github ${WORKDIR}/
COPY ./compose.yaml ${WORKDIR}/

# Markdown sources
COPY ./docs ${WORKDIR}/
COPY ./README.md ${WORKDIR}/
COPY ./LICENSE.md ${WORKDIR}/
COPY ./CODE_OF_CONDUCT.md ${WORKDIR}/

# Code source
COPY ./exercises ${WORKDIR}/exercises
COPY ./utils ${WORKDIR}/utils
COPY ./main.go ${WORKDIR}/

COPY ./go.mod ${WORKDIR}/
COPY ./go.sum ${WORKDIR}/
COPY ./Makefile ${WORKDIR}/

# markdownlint conf
COPY ./.markdownlint.yaml ${WORKDIR}/

# yamllint conf
COPY ./.yamllint ${WORKDIR}/
COPY ./.yamlignore ${WORKDIR}/
COPY ./.gitignore ${WORKDIR}/

CMD ["make", "lint"]

###############################################################################
FROM base AS development

ENV BINDIR /usr/local/bin
RUN apk add --update --no-cache make

COPY ./exercises ${WORKDIR}/exercises
COPY ./utils ${WORKDIR}/utils
COPY ./main.go ${WORKDIR}/
COPY ./go.mod ${WORKDIR}/
COPY ./go.sum ${WORKDIR}/
COPY ./Makefile ${WORKDIR}/

# CMD []
###############################################################################
FROM development AS builder

# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache ca-certificates tzdata && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

RUN make build
RUN ls -alh

# CMD []

###############################################################################
### In testing stage, can't use USER, due permissions issue
## in github actions environment:
##
## https://docs.github.com/en/actions/creating-actions/dockerfile-support-for-github-actions
##
FROM development AS testing

ENV LOG_LEVEL=INFO
ENV BRUTEFORCE=FALSE

WORKDIR /app

RUN ls -alh

CMD ["make", "test"]

###############################################################################
### In production stage
## in the production phase, "good practices" such as
## WORKDIR and USER are maintained
##
## Sources:
##   https://shakib37.medium.com/distroless-as-the-final-container-base-image-b8af961fc826
##   https://chemidy.medium.com/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

FROM scratch AS production
# FROM alpine:3.20.0 AS production

ENV WORKDIR=/app
WORKDIR ${WORKDIR}

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable
COPY --from=builder /app/bin/algorithms ${WORKDIR}/

# Use an unprivileged user.
USER appuser:appuser

# RUN ls -alh

CMD ["/app/algorithms"]
