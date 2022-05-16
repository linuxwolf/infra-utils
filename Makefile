.PHONY: help
help:
	@echo "setup ---------- : Sets up Docker buildx builder"
	@echo "caddy ---------- : Creates the Caddy image"

.PHONY: setup
setup:
	docker buildx create --driver docker-container --bootstrap --name infra-utils-builder || true

.PHONY: caddy
caddy: setup
	$(MAKE) -C caddy build