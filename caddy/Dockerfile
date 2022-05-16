FROM caddy:2.5.1-builder as builder

RUN xcaddy build \
  --with github.com/caddy-dns/digitalocean@ac6e4fd9e1352c1c179a4258a48cd03d09f095ec \
  --output /builder/caddy

FROM busybox:1.34 as caddy

COPY --from=builder /builder/caddy /usr/bin/caddy