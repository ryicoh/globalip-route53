FROM golang:1.12-alpine AS builder

WORKDIR /build
COPY . .

RUN apk add -u git && \
    go build

FROM alpine:3.10

COPY --from=builder /build/globalip-route53 /usr/local/bin/

CMD ["globalip"]
