TESTS ?= ./...

bin         := payment-reconciler
commit      := $(shell git rev-parse --short HEAD)
tag         := $(shell git tag -l 'v*-rc*' --points-at HEAD)
version     := $(shell if [[ -n "$(tag)" ]]; then echo $(tag) | sed 's/^v//'; else echo $(commit); fi)

.EXPORT_ALL_VARIABLES:
GO111MODULE = on

lint_output  := lint.txt

.PHONY: all
all: build

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build
build: fmt
	go build

.PHONY: test
test: test-unit

.PHONY: test-unit
test-unit:
	go test $(TESTS)

.PHONY: clean
clean:
	rm -f ./$(bin) ./$(bin)-*.zip $(test_path) build.log

.PHONY: package
package:
ifndef version
	$(error No version given. Aborting)
endif
	$(info Packaging version: $(version))
	$(eval tmpdir:=$(shell mktemp -d build-XXXXXXXXXX))
	cp ./$(bin) $(tmpdir)
	cp -r ./terraform  $(tmpdir)/terraform
	cd $(tmpdir) && zip -r ../$(bin)-$(version).zip $(bin) terraform
	rm -rf $(tmpdir)

.PHONY: dist
dist: clean build package

.PHONY: lint
lint: GO111MODULE=off
lint:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
	gometalinter ./... > $(lint_output); true
