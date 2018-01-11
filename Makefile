.PHONY: \
	build \
	clean \
	coveralls \
	dep \
	fmt \
	fmtcheck \
	install \
	lint \
	pretest \
	test \
	vet \

all: dep build test

SRCS = $(shell git ls-files '*.go')
PKGS = ./. ./commands ./nifcloud

build: main.go
	go build -o nifcloud-ssh-config $<

clean:
	$(foreach pkg,$(PKGS),go clean $(pkg) || exit;)

coveralls:
	goveralls -service=travis-ci

dep:
	@go get -u github.com/golang/dep/...
	@go get github.com/axw/gocov/gocov
	@go get github.com/mattn/goveralls
	dep ensure

fmt:
	gofmt -w $(SRCS)

fmtcheck:
	$(foreach file,$(SRCS),gofmt -d $(file);)

install:
	go install

lint:
	@ go get -v github.com/golang/lint/golint
	$(foreach file,$(SRCS),golint $(file) || exit;)

pretest: lint vet fmtcheck

test: pretest
	$(foreach pkg,$(PKGS),go test -v $(pkg) || exit;)

vet:
	$(foreach pkg,$(PKGS),go vet $(pkg);)
