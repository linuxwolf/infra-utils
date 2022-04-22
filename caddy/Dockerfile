FROM golang:1.17 as builder

RUN go install github.com/caddyserver/xcaddy/cmd/xcaddy@v0.2.1
RUN xcaddy build v2.4.6 \
  --with github.com/caddy-dns/digitalocean

FROM busybox:1.34 as caddy

COPY --from=builder /go/caddy /usr/local/bin/caddy
