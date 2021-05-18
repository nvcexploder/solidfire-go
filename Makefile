COMMIT := $(shell git rev-parse HEAD)
TAG := $(shell git describe --tags | sed s/v//g)
TIMESTAMP := $(shell date '+%FT%T%z')
VERSION_PKG := github.com/cloud-pi/spc-sdk-go/pkg/common/version
GOLDFLAGS := -X ${VERSION_PKG}.Timestamp=${TIMESTAMP} -X ${VERSION_PKG}.Commit=${COMMIT} -X ${VERSION_PKG}.Tag=${TAG}
GOBUILDPKGS := ./api ./
GOPRIVATE := GOPRIVATE=github.com/joyent,github.com/cloud-pi
GOLANG := 1.16
LINTER_VERSION := 1.38.0
INSTALLED_LINTER := $(shell ./bin/golangci-lint --version | sed 's/^.*[^0-9]\([0-9]*\.[0-9]*\.[0-9]*\).*$$/\1/')

PROJECT_NAME := solidfire-sdk

#ifeq (1,$(DEBUG))
#	GCFLAGS := -gcflags="all=-N -l"
#endif

# prepare to build inside of docker
GITCONFIG_JOYENT := git config --global url."git@github.com:joyent".insteadOf "https://github.com/joyent"
RSAKEY_COPY := cp -r /root/ssh /root/.ssh; chown -R root:root /root/.ssh
PREPARE_CMD := ${RSAKEY_COPY}; ${GITCONFIG_JOYENT};

all: check build

build: build-main

build-main:
	@GO111MODULE=on CGO_ENABLED=0 go build $(GCFLAGS) \
		-ldflags "-X ${VERSION_PKG}.Application=${PROJECT_NAME} ${GOLDFLAGS}" \
		-o ./bin/main ./

check:
	test -z "$(shell gofmt -l internal cmd)"
	[ -f ./bin/golangci-lint ] && [ $(INSTALLED_LINTER) = $(LINTER_VERSION) ] || \
		curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v$(LINTER_VERSION)
	./bin/golangci-lint run $(GOBUILDPKGS) --timeout 5m

fmt:
	go fmt $(GOBUILDPKGS)

test-in-docker:
	docker run --rm -v "${HOME}/.ssh":/root/ssh -v "${PWD}:/${PROJECT_NAME}" -w "/${PROJECT_NAME}" -e ${GOPRIVATE} golang:${GOLANG} \
		bash -c "${PREPARE_CMD} go get -v -u gotest.tools/gotestsum; make test-junit"

test: check
	@GO111MODULE=on go test -race -v $(GOBUILDPKGS)

# Install https://github.com/gotestyourself/gotestsum for JUnit output
test-junit: check
	gotestsum --junitfile test.xml --format short-verbose -- -race -v -coverprofile=cover.out $(GOBUILDPKGS) || \
		{ GO111MODULE=on go tool cover -html=cover.out -o test.html \
		&& exit 1 ; }
	GO111MODULE=on go tool cover -html=cover.out -o test.html
