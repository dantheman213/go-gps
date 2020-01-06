.PHONY: all build deps

all: build

build:
	CGO_ENABLED=0 \
	GO111MODULE=on \
	go build \
	-installsuffix "static" \
	-o bin/ \
	$$(find examples/*.go)

deps:
	@echo Downloading go.mod dependencies && \
		go mod download
