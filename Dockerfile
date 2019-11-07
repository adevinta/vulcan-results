FROM golang:1.13-alpine3.10 AS builder

RUN apk add --no-cache --update git dep

WORKDIR /go/src/app

COPY . .
RUN dep ensure -v
RUN go get -v ./...

FROM alpine:3.10
WORKDIR /app
COPY --from=builder /go/bin/vulcan-results .
COPY run.sh .
EXPOSE 8080
CMD ["./run.sh"]
