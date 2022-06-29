SOURCES := $(shell find . -type f -name "*.go") go.mod go.sum
BINARY := envs

.PHONY: help
help:
	@echo "build    ---------- : Build binary"
	@echo "test     ---------- : Run tests (with coverage)"
	@echo "clean    ---------- : Clean intermediate files"

$(BINARY): $(SOURCES)
	go build -o $@

.coverage:
	mkdir -p $@

.coverage/profile: .coverage $(SOURCES)
	go test \
		-coverprofile=$@ \
		./...

.coverage/index.html: .coverage/profile
	go tool cover \
		-html=$< \
		-o $@

.PHONY: build
build: $(BINARY)

.PHONY: test
test: .coverage/profile .coverage/index.html

.PHONY: clean-build
clean-build:
	rm -rf $(BINARY)

.PHONY: clean
clean: clean-build
	git clean -xfd .