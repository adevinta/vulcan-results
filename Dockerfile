FROM golang:1.13-alpine3.10 AS builder

RUN apk add --no-cache --update git dep

WORKDIR /go/src/app

COPY . .
RUN dep ensure -v
RUN go get -v ./...

FROM alpine:3.10

RUN apk add --no-cache --update gettext

WORKDIR /app
COPY --from=builder /go/bin/vulcan-results .
COPY config.toml .
COPY run.sh .

CMD ["./run.sh", "./config.toml"]
