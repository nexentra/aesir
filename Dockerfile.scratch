FROM golang:1.21.3 AS builder

RUN apt-get update -y \
 && apt-get install -y ca-certificates tzdata \
 && apt-get clean \
 && update-ca-certificates --fresh

ADD . /go/src/github/nexentra/aesir
WORKDIR /go/src/github/nexentra/aesir

RUN make build

FROM scratch

ARG BUILD_DATE
ARG BUILD_REV

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/src/github/nexentra/aesir/dist/bin/aesir /aesir

USER nobody:nogroup

ENTRYPOINT ["/aesir"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.revision="${BUILD_REV}" \
      org.opencontainers.image.title="aesir" \
      org.opencontainers.image.authors="KnockOutEZ" \
      org.opencontainers.image.vendor="KnockOutEZ" \
      org.opencontainers.image.source="https://github.com/nexentra/aesir"