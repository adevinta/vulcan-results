FROM golang:1.12.7-alpine3.10 AS builder
RUN apk add --no-cache --update git dep
RUN mkdir -p $GOPATH/src/github.com/goadesign/
WORKDIR $GOPATH/src/github.com/goadesign/
RUN git clone https://github.com/goadesign/goa.git
WORKDIR $GOPATH/src/github.com/goadesign/goa
# Pinning goa version 1.3.1
RUN git checkout fc29b77a218fb9e190849c81911ed12d25e771de
RUN go get ./...
RUN go install github.com/goadesign/goa/goagen
RUN mkdir -p $GOPATH/src/github.com/adevinta/vulcan-results
WORKDIR $GOPATH/src/github.com/adevinta/vulcan-results
COPY . .
RUN rm -rf app client tool swagger && \
  dep ensure -v && \
  rm -rf vendor && \
  goagen bootstrap -d github.com/adevinta/vulcan-results/design && \
  rm main.go && go get ./...

FROM alpine:3.10

ARG BUILD_RFC3339="1970-01-01T00:00:00Z"
ARG COMMIT="local"
ARG VERSION="dirty"

ENV BUILD_RFC3339 "$BUILD_RFC3339"
ENV COMMIT "$COMMIT"
ENV VERSION "$VERSION"

RUN apk add --no-cache --update gettext ca-certificates

WORKDIR /app
COPY --from=builder /go/bin/vulcan-results .
COPY config.toml .
COPY run.sh .
CMD ["./run.sh", "./config.toml"]
