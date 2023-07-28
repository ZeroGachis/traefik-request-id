FROM golang:1.19-bookworm

ENV GOLANGCI_LINT_VERSION=v1.53.3
ENV YAEGI_VERSION=v0.14.2
ENV CGO_ENABLED=0

# Install golangci-lint (go linter)
# Cf: https://golangci-lint.run/usage/install
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin ${GOLANGCI_LINT_VERSION}

# Install Yaegi (go interpreter used by Traefik)
RUN curl -sfL https://raw.githubusercontent.com/traefik/yaegi/master/install.sh | bash -s -- -b $(go env GOPATH)/bin ${YAEGI_VERSION}

WORKDIR /home/traefik_request_id

COPY ./ ./

RUN \
    go env -w GOPATH=/home/traefik_request_id/go && \
    go mod tidy && \
    go mod download && \
    go mod vendor
