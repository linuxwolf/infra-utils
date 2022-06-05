SOURCES := $(shell find . -type f -name "*.go")

.PHONY: help
help:
	@echo "test     ---------- : Run tests (with coverage)"
	@echo "clean    ---------- : Clean intermediate files"

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

.PHONY: test
test: .coverage/profile .coverage/index.html

.PHONY: clean
clean:
	git clean -x .