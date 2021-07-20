ARG GOLANG_VERSION=1.16.6-alpine3.14
ARG ALPINE_VERSION=3.14

FROM golang:${GOLANG_VERSION} AS builder

ADD . ${GOPATH}/src/github.com/earthquakesan/udp-server-client

WORKDIR ${GOPATH}/src/github.com/earthquakesan/udp-server-client
RUN go build -o /go/bin/udp-server udp-server.go \
 && go build -o /go/bin/udp-client udp-client.go

FROM alpine:${ALPINE_VERSION}

ENV SERVER_PORT 8866
ENV CLIENT_CONNECTION_STRING localhost:8866

COPY --from=builder /go/bin/udp-server /go/bin/udp-server
COPY --from=builder /go/bin/udp-client /go/bin/udp-client

ENTRYPOINT ["/go/bin/udp-server"]