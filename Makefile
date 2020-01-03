bin     := payment-reconciler
version := "unversioned"

lint_output  := lint.txt

.PHONY: all
all: build

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: deps
deps:
	go get ./...

.PHONY: build
build: deps fmt $(bin)

$(bin):
	go build -o ./$(bin)

.PHONY: test-deps
test-deps: deps
	go get -t ./...

.PHONY: test
test: test-unit

.PHONY: test-unit
test-unit: test-deps
	go test ./... -run 'Unit' -coverprofile=coverage.out

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
lint:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
	gometalinter ./... > $(lint_output); true