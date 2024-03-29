FROM golang:1.19.5-alpine3.17 AS golang

ARG LOC=/builds/go/src/github.com/betorvs/secretreceiver
RUN apk add --no-cache git
RUN mkdir -p $LOC
ENV GOPATH /go
COPY main.go go.mod go.sum $LOC/
COPY ./appcontext $LOC/appcontext
COPY ./config $LOC/config
COPY ./controller $LOC/controller
COPY ./domain $LOC/domain
COPY ./gateway $LOC/gateway
COPY ./tests $LOC/tests
COPY ./usecase $LOC/usecase
COPY ./utils $LOC/utils
ENV CGO_ENABLED 0
RUN cd $LOC && TESTRUN=true go test ./... && go build

FROM alpine:3.17
ARG LOC=/builds/go/src/github.com/betorvs/secretreceiver
WORKDIR /
VOLUME /tmp
RUN apk add --no-cache ca-certificates
RUN update-ca-certificates
RUN mkdir -p /app
RUN addgroup -g 1000 -S app && \
    adduser -u 1000 -G app -S -D -h /app app && \
    chmod 755 /app
COPY --from=golang $LOC/secretreceiver /app

EXPOSE 8080
RUN chmod +x /app/secretreceiver
WORKDIR /app    
USER app
CMD ["/app/secretreceiver"]
