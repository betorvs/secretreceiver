FROM golang:1.13.6-alpine3.11 AS golang

RUN apk add --no-cache git
RUN mkdir -p /builds/go/src/github.com/betorvs/secretreceiver/
ENV GOPATH /builds/go
COPY . /builds/go/src/github.com/betorvs/secretreceiver/
ENV CGO_ENABLED 0
RUN cd /builds/go/src/github.com/betorvs/secretreceiver/ && go build

FROM alpine:3.11
WORKDIR /
VOLUME /tmp
RUN apk add --no-cache ca-certificates
COPY --from=golang /builds/go/src/github.com/betorvs/secretreceiver/secretreceiver /
RUN update-ca-certificates

EXPOSE 9090
RUN chmod +x /secretreceiver
CMD ["/secretreceiver"]