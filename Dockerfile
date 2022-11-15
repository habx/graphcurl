FROM alpine:3.16.3

ARG CREATED
ARG REVISION
ARG VERSION
ARG TITLE
ARG SOURCE
ARG AUTHORS
LABEL org.opencontainers.image.created=$CREATED \
        org.opencontainers.image.revision=$REVISION \
        org.opencontainers.image.title=$TITLE \
        org.opencontainers.image.source=$SOURCE \
        org.opencontainers.image.version=$VERSION \
        org.opencontainers.image.authors=$AUTHORS \
        org.opencontainers.image.vendor="Habx"

ENV TZ=Europe/Paris
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /go/src/github.com/habx/graphcurl/

COPY dist/graphcurl_linux_amd64/graphcurl_linux_amd64 /go/src/github.com/habx/graphcurl/graphcurl_linux_amd64

ENTRYPOINT ["./graphcurl_linux_amd64"]
