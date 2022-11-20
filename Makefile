ifndef BUILDHASH
BUILDHASH ?= $(shell git log -1 --pretty='format:%h')
endif

BUILDNUMBER ?= $(shell date +%Y%m%d.%H%M-%S)
VERSION ?= $(BUILDHASH)
LDFLAGS = "-X main.version=$(VERSION) -X main.buildNumber=$(BUILDNUMBER)"

all: bin/server

bin/server:
	GOSUMDB=off go build -mod=vendor -ldflags $(LDFLAGS) -v -o bin/server ./cmd/server

vendor:
	GOSUMDB=off go mod vendor

clean:
	rm -fv bin/server
