FROM golang:alpine as builder

WORKDIR /go/src/github.com/habx/graphcurl
ADD dist/graphcurl_linux_amd64.gz /go/src/github.com/habx/graphcurl

RUN apk add --no-cache ca-certificates tzdata && \
    gzip -d /go/src/github.com/habx/graphcurl/graphcurl_linux_amd64.gz && \
    chmod +x /go/src/github.com/habx/graphcurl/graphcurl_linux_amd64

FROM alpine:edge

ENV TZ=Europe/Paris

WORKDIR /go/src/github.com/habx/graphcurl/

RUN apk add --no-cache jq

COPY --from=builder /usr/share/zoneinfo/ /usr/share/zoneinfo/
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=builder /go/src/github.com/habx/graphcurl/graphcurl_linux_amd64 .

ENTRYPOINT ["./graphcurl_linux_amd64"]
