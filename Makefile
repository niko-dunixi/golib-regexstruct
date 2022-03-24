GO := go

.PHONY: all
all: test-all examples-all

.PHONY: test-all
test-all:
	$(MAKE) GO=go1.17 test
	$(MAKE) GO=go1.18 test

.PHONY: test
test:
	$(GO) test -v ./...


.PHONY: examples-all
examples-all:
	$(MAKE) examples-gorillamux

.PHONY: examples-gorillamux
examples-gorillamux:
	$(MAKE) -C ./examples/gorillamux GO=$(GO) test
