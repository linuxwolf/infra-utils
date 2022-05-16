.PHONY: build
build:
	docker buildx --builder infra-utils-builder \
		build \
		--cache-from type=local,src=../.cache \
		--cache-to type=local,mode=max,dest=../.cache \
		--tag linuxwolf/caddy:latest \
		--load \
		.